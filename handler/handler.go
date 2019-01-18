package handler

import (
    "encoding/json"
    "gonum.org/v1/gonum/mat"
    "net/http"
    "net/url"
    "math"
)

const PARSING_ERROR = "parsingError"
const VALIDATION_ERROR = "validationError"

type requestBody struct {
    Degree int `json:"degree"`
    Value []float64 `json:"value"`
}

type responseBody struct {
    Value []float64 `json:"value"`
}

type unitaryNDilation func(*mat.Dense, int) (*mat.Dense, error)

func validateRequestBody(b requestBody) url.Values {
    errs := url.Values{}

    if b.Degree == 0 {
        errs.Add("degree", "degree must be an integer greater than zero")
    }

    n := int(math.Sqrt(float64(len(b.Value))))
    if len(b.Value) == 0 || int(math.Pow(float64(n), 2)) != len(b.Value) {
        errs.Add("value", "value must contain a square number greater than zero of numbers")
    }
    return errs
}

func sendBadRequestResponse(w http.ResponseWriter, errorType string, errs url.Values) {
    err := map[string]interface{}{errorType: errs}
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(err)
}

func denseToSlice(u *mat.Dense) (data []float64) {
    m, _ := u.Dims()
    for i := 0; i < m; i++ {
        raw := u.RawRowView(i)
        data = append(data, raw...)
    }
    return
}

func handleDilationPost(dilation unitaryNDilation, w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var b requestBody
    err := decoder.Decode(&b)
    if err != nil {
        errs := url.Values{}
        errs.Add("body", "json malformed")
        sendBadRequestResponse(w, PARSING_ERROR, errs)
        return
    }

    errs := validateRequestBody(b)
    if len(errs) > 0 {
        sendBadRequestResponse(w, VALIDATION_ERROR, errs)
        return
    }

    n := int(math.Sqrt(float64(len(b.Value))))
    t := mat.NewDense(n, n, b.Value)
    unitary, e := dilation(t, b.Degree)

    if e != nil {
        errs.Add("value", "value must represent a real contraction")
        sendBadRequestResponse(w, VALIDATION_ERROR, errs)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(responseBody{ Value: denseToSlice(unitary) })
}

func DilationHandler(dilation unitaryNDilation) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case http.MethodPost:
                handleDilationPost(dilation, w, r)
            default:
                http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
            }
    }
}