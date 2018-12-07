package blockMatrix

import (
    "gonum.org/v1/gonum/mat"
)

func NewBlockMatrix(rows [][]*mat.Dense) (*mat.Dense, bool) {
    var m, n int
    m, n = 0, 0

    for i, row := range rows {
        k, _ := row[0].Dims()
        m += k

        if (i == 0) {
            for _, matrix := range row {
                _, l := matrix.Dims()
                n += l
            }
        }
    }

    return mat.NewDense(m, n, nil), true
}
