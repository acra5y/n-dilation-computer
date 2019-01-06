package main

import (
    "gonum.org/v1/gonum/mat"
    "math/rand"
    "fmt"
    "github.com/acra5y/n-dilation-computer/positiveDefinite"
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

type PositiveDefiniteResult struct {
    isPositiveDefinite bool
    matrix mat.Matrix
}

func positivedefiniteMatrix() *mat.Dense {
    data := []float64{1,2,2,100}
    return mat.NewDense(2, 2, data)
}

func isPositiveDefinite(a *mat.Dense, c chan PositiveDefiniteResult) {
    m, n := a.Dims()
    v := mat.NewDense(m, n, nil)
    v.Clone(a)
    eigen := mat.Eigen{}
    candidate := positiveDefinite.PositiveDefiniteCandidate{Value: a}
    result, _ := candidate.IsPositiveDefinite(&eigen)

    c <- PositiveDefiniteResult{isPositiveDefinite: result, matrix: candidate.Value}
}

func main() {
    dimension := 6
    data := make([]float64, 36)
    for i := range data {
        data[i] = rand.NormFloat64()
    }
    a := mat.NewDense(dimension, dimension, data)

    printDivider()
    printM(a)

    values := mat.Eigen{}
    values.Factorize(a, false, true)

    c := make(chan PositiveDefiniteResult, 2)
    go isPositiveDefinite(a, c)
    go isPositiveDefinite(positivedefiniteMatrix(), c)

    var x, y PositiveDefiniteResult
    x = <- c
    y = <- c
    fmt.Printf("a=%v %v %v\n", x.matrix, "is positivedefinite: ", x.isPositiveDefinite)
    printDivider()
    fmt.Printf("a=%v %v %v\n", y.matrix, "is positivedefinite: ", y.isPositiveDefinite)
    printDivider()

    if x.isPositiveDefinite {
        m, n := x.matrix.Dims()
        v := mat.NewDense(m, n, nil)
        v.Clone(x.matrix)

        res, _ := squareRoot.Calculate(v)
        prod := mat.NewDense(m, n, nil)
        prod.Product(res, res)
        fmt.Printf("M=%v sq=%v sq^2=%v\n", v, res, prod)
        printDivider()
    }

    if y.isPositiveDefinite {
        m, n := y.matrix.Dims()
        v := mat.NewDense(m, n, nil)
        v.Clone(y.matrix)

        res, _ := squareRoot.Calculate(v)
        prod := mat.NewDense(m, n, nil)
        prod.Product(res, res)
        fmt.Printf("M=%v sq=%v sq^2=%v\n", v, res, prod)
        printDivider()
    }

    fmt.Println("done")
}
