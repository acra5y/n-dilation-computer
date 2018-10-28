package ndilation

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestUnitaryNDilation(t *testing.T) {
    tables := []struct {
    	value *mat.Dense
	}{
		{value: mat.NewDense(2, 2, nil)},
	}

	for _, table := range tables {
		ndilation := Dilation{ N: 1 }
		err := ndilation.unitaryNDilation(table.value)

		if err != nil {
			t.Errorf("Unexpected err, want: %v, got: %v", nil, err)
		}

		unitary := ndilation.Value()

		if !mat.Equal(unitary, table.value) {
			t.Errorf("Wrong matrix returned, got: %v, want: %v", mat.NewDense(2, 2, nil), unitary)
		}
	}
}