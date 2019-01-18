package main

import (
    "github.com/acra5y/n-dilation-computer/dilation"
    "github.com/acra5y/n-dilation-computer/handler"
    "github.com/acra5y/n-dilation-computer/positiveDefinite"
    "github.com/acra5y/n-dilation-computer/squareRoot"
    "github.com/acra5y/n-dilation-computer/blockMatrix"
    "gonum.org/v1/gonum/mat"
    "log"
    "net/http"
)

func unitaryNDilation(t *mat.Dense, n int) (*mat.Dense, error) {
    return dilation.UnitaryNDilation(positiveDefinite.IsPositiveDefinite, squareRoot.Calculate, blockMatrix.NewBlockMatrixFromSquares, t, n)
}

func main() {
    http.HandleFunc("/dilation", handler.DilationHandler(unitaryNDilation))
    port := ":8080"
    log.Printf("Listening on port %s...", port)
    log.Fatal(http.ListenAndServe(port, nil))
}
