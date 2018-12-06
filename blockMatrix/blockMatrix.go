package blockMatrix

import (
    "gonum.org/v1/gonum/mat"
)

func NewBlockMatrix([][]*mat.Dense) *mat.Dense {
	return mat.NewDense(0, 0, nil)
}