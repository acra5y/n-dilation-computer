package blockMatrix

import (
    "gonum.org/v1/gonum/mat"
    "testing"
)

var dummyMatrix = mat.NewDense(2, 2, nil)

func createBlocks() (blocks [][]*mat.Dense) {
    blocks = make([][]*mat.Dense, 1)
    firstRow := make([]*mat.Dense, 1)
    firstRow[0] = mat.NewDense(0, 0, nil)
    blocks[0] = firstRow
    return
}

func TestNewBlockMatrix(t *testing.T) {
    tables := []struct {
        desc string
        blocks [][]*mat.Dense
        expected *mat.Dense
    }{
        {blocks: createBlocks(), expected: mat.NewDense(0, 0, nil), desc: "returns correct matrix"},
    }

    for _, table := range tables {
        t.Run(table.desc, func(t *testing.T) {
            t.Parallel()

            blockMatrix := NewBlockMatrix(table.blocks)

            if !mat.Equal(blockMatrix, table.expected) {
                t.Errorf("NewBlockMatrix returned wrong value, got: %v, want: %v.", blockMatrix, table.expected)
            }
        })
    }
}
