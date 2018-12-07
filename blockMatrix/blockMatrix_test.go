package blockMatrix

import (
    "gonum.org/v1/gonum/mat"
    "testing"
)

func createRows() (rows [][]*mat.Dense) {
    rows = make([][]*mat.Dense, 2)

    for i := range rows {
        row := make([]*mat.Dense, 3)

        for j := range row {
            row[j] = mat.NewDense(1, 3, nil)
        }

        rows[i] = row
    }
    return
}

func TestNewBlockMatrix(t *testing.T) {
    tables := []struct {
        desc string
        rows [][]*mat.Dense
        expected *mat.Dense
        ok bool
    }{
        {rows: createRows(), expected: mat.NewDense(2, 9, nil), desc: "returns correct matrix", ok: true},
    }

    for _, table := range tables {
        t.Run(table.desc, func(t *testing.T) {
            t.Parallel()

            blockMatrix, ok := NewBlockMatrix(table.rows)

            if ok !=table.ok {
                t.Errorf("NewBlockMatrix returned wrong value for ok, got: %t, want: %t.", ok, table.ok)
            }

            if !mat.Equal(blockMatrix, table.expected) {
                t.Errorf("NewBlockMatrix returned wrong value, got: %v, want: %v.", blockMatrix, table.expected)
            }
        })
    }
}
