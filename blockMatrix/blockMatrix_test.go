package blockMatrix

import (
    "gonum.org/v1/gonum/mat"
    "testing"
)

var dummyMatrix = mat.NewDense(2, 2, nil)

func createRows() (rows [][]*mat.Dense) {
    rows = make([][]*mat.Dense, 2)

    for i := range rows {
        row := make([]*mat.Dense, 1)
        row[0] = mat.NewDense(1, 3, nil)
        rows[i] = row
    }
    return
}

func TestNewBlockMatrix(t *testing.T) {
    tables := []struct {
        desc string
        rows [][]*mat.Dense
        expected *mat.Dense
    }{
        {rows: createRows(), expected: mat.NewDense(2, 3, nil), desc: "returns correct matrix"},
    }

    for _, table := range tables {
        t.Run(table.desc, func(t *testing.T) {
            t.Parallel()

            blockMatrix := NewBlockMatrix(table.rows)

            if !mat.Equal(blockMatrix, table.expected) {
                t.Errorf("NewBlockMatrix returned wrong value, got: %v, want: %v.", blockMatrix, table.expected)
            }
        })
    }
}
