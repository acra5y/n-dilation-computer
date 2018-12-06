package blockMatrix

import (
    "gonum.org/v1/gonum/mat"
)

func NewBlockMatrix(rows [][]*mat.Dense) *mat.Dense {
    var m, n int
    m, n = 0, 0

    for _, row := range rows {
        k, _ := row[0].Dims()
        m += k
    }

    _, n = rows[0][0].Dims()

    return mat.NewDense(m, n, nil)
}
