package dilation

import (
    "github.com/acra5y/n-dilation-computer/positiveDefinite"
    "gonum.org/v1/gonum/mat"
    "testing"
)

func testIsPositiveDefinite(t *testing.T, expected []*mat.Dense) isPositiveDefinite {
    calls := 0
    return func(a positiveDefinite.EigenComputer, candidate *mat.Dense) (bool, error) {
        if !mat.Equal(expected[calls], candidate) {
            t.Errorf("Unexpected argument in call to testIsPositiveDefinite. Got %v: ,want: %v", candidate, expected[calls])
        }
        calls++
        return true, nil
    }
}

func testSquareRoot(t *testing.T, expected []*mat.Dense) squareRoot {
    calls := 0
    return func(a *mat.Dense) (*mat.Dense, error) {
        if !mat.Equal(expected[calls], a) {
            t.Errorf("Unexpected argument in call %d to squareRoot. Got: %v, want: %v", calls + 1, a, expected[calls])
        }
        calls++
        return mat.NewDense(2, 2, nil), nil
    }
}

func testNewBlockMatrixFromSquares(t *testing.T, expected [][]*mat.Dense) newBlockMatrixFromSquares {
    return func(rows [][]*mat.Dense) (*mat.Dense, error) {
        for i, row := range rows {
            for j, m := range row {
                if !mat.Equal(expected[i][j], m) {
                    t.Errorf("Unexpected argument in call to newBlockMatrixFromSquares. Wrong block at position (%d, %d), got : %v, want: %v", i, j, m, expected[i][j])
                }
            }
        }
        return mat.NewDense(2, 2, nil), nil
    }
}

func TestUnitaryNDilation(t *testing.T) {
    tables := []struct {
        value *mat.Dense
        expectedInSqrt []*mat.Dense
        expectedRows [][]*mat.Dense
    }{
        {
            value: mat.NewDense(2, 2, nil),
            expectedInSqrt: []*mat.Dense{mat.NewDense(2, 2, []float64{1,0,0,1,}),mat.NewDense(2, 2, []float64{1,0,0,1,}),},
            expectedRows: [][]*mat.Dense{
                []*mat.Dense{mat.NewDense(2, 2, nil),mat.NewDense(2, 2, nil),},
                []*mat.Dense{mat.NewDense(2, 2, nil),mat.NewDense(2, 2, nil),},
            },
        },
        {
            value: mat.NewDense(2, 2, []float64{0.5,0.5,0,0.5,}),
            expectedInSqrt: []*mat.Dense{mat.NewDense(2, 2, []float64{0.5,-0.25,-0.25,0.75,}),mat.NewDense(2, 2, []float64{0.75,-0.25,-0.25,0.5,}),},
            expectedRows: [][]*mat.Dense{
                []*mat.Dense{mat.NewDense(2, 2, []float64{0.5,0.5,0,0.5,}),mat.NewDense(2, 2, nil),},
                []*mat.Dense{mat.NewDense(2, 2, nil),mat.NewDense(2, 2, []float64{-0.5,0,-0.5,-0.5}),},
            },
        },
    }

    for _, table := range tables {
        unitary, err := UnitaryNDilation(
            testIsPositiveDefinite(t, table.expectedInSqrt),
            testSquareRoot(t, table.expectedInSqrt),
            testNewBlockMatrixFromSquares(t, table.expectedRows),
            table.value,
        )

        if err != nil {
            t.Errorf("Unexpected err, want: %v, got: %v", nil, err)
        }

        if !mat.Equal(unitary, mat.NewDense(2, 2, nil)) {
            t.Errorf("Wrong matrix returned, want: %v, got: %v", mat.NewDense(2, 2, nil), unitary)
        }
    }
}
