package squareRoot

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
)

type SquareRoot struct {
    Value mat.Matrix
}

func (squareRoot *SquareRoot) Calculate() (mat.Matrix, error) {
    return mat.NewDense(0, 0, nil), fmt.Errorf("Not Implemented")
}
