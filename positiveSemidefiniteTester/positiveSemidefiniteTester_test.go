package positiveSemidefiniteTester

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestPsd(t *testing.T) {
    tables := []struct {
		matrix mat.Matrix
		isPsd bool
	}{
		{matrix: mat.NewDense(2, 2, []float64{1,2,2,100}), isPsd: true},
		{matrix: mat.NewDense(2, 2, []float64{1,2,2,-100}), isPsd: false},
		{matrix: mat.NewDense(2, 2, []float64{1,2,3,4}), isPsd: false},

	}

	for _, table := range tables {
	candidate := PositiveSemidefiniteCandidate{ Value: table.matrix }
    isPsd := candidate.IsPositiveSemidefinite()
		if isPsd != table.isPsd {
			t.Errorf(
				"IsPositiveSemidefinite of (\n%v\n) was incorrect, got: %t, want: %t.",
				mat.Formatted(table.matrix, mat.Prefix("    "), mat.Squeeze()),
				isPsd,
				table.isPsd,
			)
		}
	}
}
