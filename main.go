package main

import (
    "gonum.org/v1/gonum/mat"
    "math/rand"
    "fmt"
    "github.com/acra5y/n-dilation-computer/positiveSemidefinite"
    "github.com/acra5y/n-dilation-computer/squareRoot"
)

func printDivider() {
    fmt.Printf("%s\n", "--------------------------------")
}

func printM(a mat.Matrix) {
    fa := mat.Formatted(a, mat.Prefix("    "), mat.Squeeze())

    fmt.Printf("\na = %v\n\n", fa)
    printDivider()
}

type PsdResult struct {
    isPsd bool
    matrix mat.Matrix
}

func psdMatrix() *mat.Dense {
    data := []float64{1,2,2,100}
    return mat.NewDense(2, 2, data)
}

func isPsd(a *mat.Dense, c chan PsdResult) {
    m, n := a.Dims()
    v := mat.NewDense(m, n, nil)
    v.Clone(a)
    eigen := mat.Eigen{}
    candidate := positiveSemidefinite.PositiveSemidefiniteCandidate{Value: a}
    result, _ := candidate.IsPositiveSemidefinite(&eigen)

    c <- PsdResult{isPsd: result, matrix: candidate.Value}
}

func main() {
    dimension := 6
    data := make([]float64, 36)
    for i := range data {
        data[i] = rand.NormFloat64()
    }
    a := mat.NewDense(dimension, dimension, data)

    tr := mat.Trace(a)

    fmt.Printf("%g\n", tr)
    printDivider()
    printM(a)
    printDivider()

    values := mat.Eigen{}
    values.Factorize(a, false, true)

    printDivider()
    for _, v := range values.Values(nil) {
        fmt.Printf("%g\n", v)
    }

    var eigen mat.Eigen
    eigen.Factorize(a, false, false)

    for _, v := range eigen.Values(nil) {
        fmt.Printf("%g\n", v)
    }

    printDivider()
    c := make(chan PsdResult, 2)
    go isPsd(a, c)
    go isPsd(psdMatrix(), c)

    var x, y PsdResult
    x = <- c
    y = <- c
    fmt.Printf("a=%v %v %v\n", x.matrix, "is psd: ", x.isPsd)
    printDivider()
    fmt.Printf("a=%v %v %v\n", y.matrix, "is psd: ", y.isPsd)
    printDivider()

    if x.isPsd {
        m, n := x.matrix.Dims()
        v := mat.NewDense(m, n, nil)
        v.Clone(x.matrix)

        sqr := squareRoot.SquareRoot{ C: v }
        res, _ := sqr.Calculate()
        prod := mat.NewDense(m, n, nil)
        prod.Product(res, res)
        fmt.Printf("M=%v sq=%v sq^2=%v\n", v, res, prod)
        printDivider()
    }

    if y.isPsd {
        m, n := y.matrix.Dims()
        v := mat.NewDense(m, n, nil)
        v.Clone(y.matrix)

        sqr := squareRoot.SquareRoot{ C: v }
        res, _ := sqr.Calculate()
        prod := mat.NewDense(m, n, nil)
        prod.Product(res, res)
        fmt.Printf("M=%v sq=%v sq^2=%v\n", v, res, prod)
        printDivider()
    }

    fmt.Println("done")
}
