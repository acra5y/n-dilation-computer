package blockMatrix

import (
    "math"
    "gonum.org/v1/gonum/mat"
)

func NewBlockMatrixFromSquares(rows [][]*mat.Dense) (*mat.Dense, bool) {
    var d0, d int
    d0, _ = rows[0][0].Dims()
    d = d0 * len(rows)

    var data []float64
    data = make([]float64, int(math.Pow(float64(d), 2)))

    for i, _ := range rows {
        for j := 0; j < d0; j++ {
            for k, matrix := range rows[i] {
                raw := matrix.RawRowView(j)
                offsetHandledBlockRows := len(rows[i]) * d
                offsetCurrentBlocks := j * len(rows[i]) * d0
                offsetHandledBlocks := k * d0
                index := i * offsetHandledBlockRows + offsetCurrentBlocks + offsetHandledBlocks

                data = append(data[:index], append(raw, data[(index + d0):]...)...)
            }
        }
    }

    return mat.NewDense(d, d, data), true
}
