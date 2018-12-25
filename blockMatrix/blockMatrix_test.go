package blockMatrix

import (
    "gonum.org/v1/gonum/mat"
    "testing"
)

func createRows() (rows [][]*mat.Dense) {
    rows = make([][]*mat.Dense, 2)

    for i := range rows {
        row := make([]*mat.Dense, 2)

        for j := range row {
        value := float64(2 * i + j)
            data := make([]float64, 4)
            for k, _ := range data {
                data[k] = value
            }
            row[j] = mat.NewDense(2, 2, data)
        }

        rows[i] = row
    }
    return
}

func TestNewBlockMatrixFromSquares(t *testing.T) {
    tables := []struct {
        desc string
        rows [][]*mat.Dense
        expected *mat.Dense
        ok bool
    }{
        {rows: createRows(), expected: mat.NewDense(4, 4, []float64{0,0,1,1,0,0,1,1,2,2,3,3,2,2,3,3}), desc: "returns correct matrix", ok: true},
    }

    for _, table := range tables {
        t.Run(table.desc, func(t *testing.T) {
            t.Parallel()

            blockMatrix, ok := NewBlockMatrixFromSquares(table.rows)

            if ok !=table.ok {
                t.Errorf("NewBlockMatrix returned wrong value for ok, got: %t, want: %t.", ok, table.ok)
            }

            if !mat.Equal(blockMatrix, table.expected) {
                t.Errorf("NewBlockMatrix returned wrong value, got: %v, want: %v.", blockMatrix, table.expected)
            }
        })
    }
}
