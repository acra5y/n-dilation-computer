package dilation

import (
    "github.com/acra5y/n-dilation-computer/positiveDefinite"
    "gonum.org/v1/gonum/mat"
    "testing"
)

func isPositiveDefiniteMock(a positiveDefinite.EigenComputer, candidate *mat.Dense) (bool, error) {
    return true, nil
}

func testSquareRoot(t *testing.T, expected []*mat.Dense) squareRoot {
    calls := 0
    return func(a *mat.Dense) (*mat.Dense, error) {
        if !mat.Equal(expected[calls], a) {
            t.Errorf("Unexpected argument in call %d to squareRoot. Got: %v, want: %v", calls + 1, a, expected)
        }
        calls++
        return mat.NewDense(2, 2, nil), nil
    }
}

func newBlockMatrixFromSquaresMock(a [][]*mat.Dense) (*mat.Dense, error) {
    return mat.NewDense(2, 2, nil), nil
}

func TestUnitaryNDilation(t *testing.T) {
    tables := []struct {
        value *mat.Dense
        expectedInSqrt []*mat.Dense
    }{
        {value: mat.NewDense(2, 2, nil), expectedInSqrt: []*mat.Dense{mat.NewDense(2, 2, []float64{1,0,0,1,}),mat.NewDense(2, 2, []float64{1,0,0,1,}),},},
    }

    for _, table := range tables {
        unitary, err := UnitaryNDilation(isPositiveDefiniteMock, testSquareRoot(t, table.expectedInSqrt), newBlockMatrixFromSquaresMock, table.value)

        if err != nil {
            t.Errorf("Unexpected err, want: %v, got: %v", nil, err)
        }

        if !mat.Equal(unitary, table.value) {
            t.Errorf("Wrong matrix returned, want: %v, got: %v", mat.NewDense(4, 4, nil), unitary)
        }
    }
}
