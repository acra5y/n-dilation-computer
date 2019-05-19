package handler

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func testUnitaryNDilation(test *testing.T, expectedT *mat.Dense, expectedN int, errorToThrow error) unitaryNDilation {
    return func(t *mat.Dense, n int) (*mat.Dense, error) {
        if !mat.Equal(t, expectedT) {
            test.Errorf("unexpected argument t in call to unitaryNDilation. Got: %v, want: %v", t, expectedT)
        }

        if n != expectedN {
            test.Errorf("unexpected argument n in call to unitaryNDilation. Got: %d, want: %d", n, expectedN)
        }

        return mat.NewDense(2, 2, nil), errorToThrow
    }
}

func TestDilationHandlerPost(t *testing.T) {
    tables := []struct {
        desc string
        body string
        errorToThrow error
        expectedT *mat.Dense
        expectedN int
        expectedStatus int
        expectedBody string
    }{
        {
            desc: "returns 200",
            body: `{"degree":2,"value":[0,0,0,0]}`,
            expectedT: mat.NewDense(2, 2, nil),
            expectedN: 2,
            expectedStatus: http.StatusOK,
            expectedBody: "{\"value\":[0,0,0,0]}\n",
        },
        {
            desc: "returns 400 for invalid json",
            body: `{"degree":}`,
            expectedStatus: http.StatusBadRequest,
            expectedBody: "{\"parsingError\":{\"body\":[\"json malformed\"]}}\n",
        },
        {
            desc: "returns 400 if degree is missing",
            body: `{"value":[0,0,0,0]}`,
            expectedStatus: http.StatusBadRequest,
            expectedBody: "{\"validationError\":{\"degree\":[\"degree must be an integer greater than zero\"]}}\n",
        },
        {
            desc: "returns 400 if value is missing",
            body: `{"degree":2}`,
            expectedStatus: http.StatusBadRequest,
            expectedBody: "{\"validationError\":{\"value\":[\"value must contain a square number greater than zero of numbers\"]}}\n",
        },
        {
            desc: "returns 400 if value and degree are missing",
            body: `{"foo":"bar"}`,
            expectedStatus: http.StatusBadRequest,
            expectedBody: "{\"validationError\":{\"degree\":[\"degree must be an integer greater than zero\"],\"value\":[\"value must contain a square number greater than zero of numbers\"]}}\n",
        },
        {
            desc: "returns 400 if unitaryNDilation returns error",
            body: `{"degree":2,"value":[0,0,0,0]}`,
            errorToThrow: fmt.Errorf("test-error"),
            expectedT: mat.NewDense(2, 2, nil),
            expectedN: 2,
            expectedStatus: http.StatusBadRequest,
            expectedBody: "{\"validationError\":{\"value\":[\"value must represent a real contraction\"]}}\n",
        },
    }

    for _, table := range tables {
        table := table
        t.Run(table.desc, func(t *testing.T) {
            t.Parallel()
            reader := strings.NewReader(table.body)
            req, err := http.NewRequest("POST", "/mock", reader)
            req.Header.Set("Content-Type", "application/json")
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            handler := http.HandlerFunc(DilationHandler(testUnitaryNDilation(t, table.expectedT, table.expectedN, table.errorToThrow)))
            handler.ServeHTTP(rr, req)

            if status := rr.Code; status != table.expectedStatus {
                t.Errorf("handler returned wrong status code: got %v want %v", status, table.expectedStatus)
            }

            if r := rr.Body.String(); r != table.expectedBody {
                t.Errorf("handler returned unexpected body: got %v want %v", r, table.expectedBody)
            }
        })
    }
}


func TestDilationHandlerOptions(t *testing.T) {
    req, err := http.NewRequest("OPTIONS", "/mock", nil)
    req.Header.Set("Content-Type", "application/json")
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(DilationHandler(func(t *mat.Dense, n int) (*mat.Dense, error) { return mat.NewDense(0, 0, nil), nil }))
    handler.ServeHTTP(rr, req)
    expectedStatus := http.StatusOK

    if status := rr.Code; status != expectedStatus {
        t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
    }

    headers := rr.Header()
    expectedHeaders := make(map[string]string)
    expectedHeaders["Access-Control-Allow-Methods"] = "POST,OPTIONS"
    expectedHeaders["Content-Type"] = "application/json"

    for header, expectedValue := range expectedHeaders {
        if res := headers.Get(header); res != expectedValue {
            t.Errorf("handler returned wrong %s header: got %v want %v", header, res, expectedValue)
        }
    }
}
