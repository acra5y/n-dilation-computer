package blockMatrix

import (
    "math"
    "gonum.org/v1/gonum/mat"
)

func NewBlockMatrixFromSquares(rows [][]*mat.Dense) (*mat.Dense, bool) {
    var d0, d, entriesPerBlock, entriesPerBlockRow, blockDim int
    d0, _ = rows[0][0].Dims()
    entriesPerBlock = int(math.Pow(float64(d0), 2))
    blockDim = len(rows)
    entriesPerBlockRow = entriesPerBlock * blockDim

    d = d0 * blockDim

    var data []float64
    data = make([]float64, int(math.Pow(float64(d), 2)))

    for i, _ := range rows {
        offsetHandledBlockRows := i * entriesPerBlockRow
        for j := 0; j < d0; j++ {
            offsetCurrentBlocks := j * d
            for k, matrix := range rows[i] {
                raw := matrix.RawRowView(j)
                offsetHandledBlocks := k * d0
                index := offsetHandledBlockRows + offsetCurrentBlocks + offsetHandledBlocks

                data = append(data[:index], append(raw, data[(index + d0):]...)...)
            }
        }
    }

    return mat.NewDense(d, d, data), true
}
