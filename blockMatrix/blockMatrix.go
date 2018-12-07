package blockMatrix

import (
    "math"
    "gonum.org/v1/gonum/mat"
)

func NewBlockMatrixFromSquares(rows [][]*mat.Dense) (*mat.Dense, bool) {
    var d0, d int
    d0, _ = rows[0][0].Dims()
    d = int(math.Pow(float64(d0), 2))

    return mat.NewDense(d, d, nil), true
}
