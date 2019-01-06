package ndilation

import (
	"github.com/acra5y/n-dilation-computer/positiveDefinite"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func isPositiveDefiniteMock(a positiveDefinite.EigenComputer) (bool, error) {
	return true, nil
}

func squareRootMock(a *mat.Dense) (*mat.Dense, error) {
	return mat.NewDense(1, 1, nil), nil
}

func newBlockMatrixFromSquaresMock(a [][]*mat.Dense) (*mat.Dense, error) {
	return mat.NewDense(1, 1, nil), nil
}

func TestUnitaryNDilation(t *testing.T) {
    tables := []struct {
    	value *mat.Dense
	}{
		{value: mat.NewDense(2, 2, nil)},
	}

	for _, table := range tables {
		ndilation := Dilation{ N: 1 }
		err := ndilation.UnitaryNDilation(isPositiveDefiniteMock, squareRootMock, newBlockMatrixFromSquaresMock, table.value)

		if err != nil {
			t.Errorf("Unexpected err, want: %v, got: %v", nil, err)
		}

		unitary := ndilation.Value()

		if !mat.Equal(unitary, table.value) {
			t.Errorf("Wrong matrix returned, got: %v, want: %v", mat.NewDense(2, 2, nil), unitary)
		}
	}
}