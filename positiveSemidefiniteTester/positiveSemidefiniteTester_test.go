package positiveSemidefiniteTester

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

type EigenMock struct {
	MockData []complex128
}

func (eigen EigenMock) Factorize(a mat.Matrix, left, right bool) bool {
	return true
}

func (eigen EigenMock) Values(dst []complex128) []complex128 {
	return eigen.MockData
}

func createEigenMock(values []complex128) EigenMock {
	eigen := EigenMock{ MockData: values }
	return eigen
}

var dummyMatrix = mat.NewDense(2, 2, []float64{0,0,0,0})

func TestPsdForSymmetricMatrix(t *testing.T) {
    tables := []struct {
		values []complex128
		isPsd bool
	}{
		{values: []complex128{0,5}, isPsd: true},
		{values: []complex128{0,-5}, isPsd: false},
		{values: []complex128{0,complex(5, 7)}, isPsd: false},
		//{matrix: mat.NewDense(2, 2, []float64{1,2,3,4}), isPsd: false},
	}

	for _, table := range tables {
	candidate := PositiveSemidefiniteCandidate{ Value: dummyMatrix }
    isPsd := candidate.IsPositiveSemidefinite(createEigenMock(table.values))
		if isPsd != table.isPsd {
			t.Errorf("IsPositiveSemidefinite was incorrect, got: %t, want: %t.", isPsd, table.isPsd)
		}
	}
}

func TestPsdForNotSymmetricMatrix(t *testing.T) {
	candidate := PositiveSemidefiniteCandidate{ Value: mat.NewDense(2, 2, []float64{0,1,0,0}) }
	isPsd := candidate.IsPositiveSemidefinite(createEigenMock([]complex128{0,5}))
	if isPsd != false {
		t.Errorf("IsPositiveSemidefinite was incorrect, got: %t, want: %t.", isPsd, false)
	}
}
