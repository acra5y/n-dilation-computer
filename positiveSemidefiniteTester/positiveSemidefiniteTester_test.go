package positiveSemidefiniteTester

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"reflect"
	"testing"
)

type EigenMock struct {
	factorizeOk bool
	mockData []complex128
}

func (eigen EigenMock) Factorize(a mat.Matrix, left, right bool) bool {
	return eigen.factorizeOk
}

func (eigen EigenMock) Values(dst []complex128) []complex128 {
	return eigen.mockData
}

func createEigenMock(factorizeOk bool, values []complex128) EigenMock {
	eigen := EigenMock{ mockData: values, factorizeOk: factorizeOk }
	return eigen
}

var dummyMatrix = mat.NewDense(2, 2, []float64{0,0,0,0})

func TestPsdForMatrix(t *testing.T) {
    tables := []struct {
    	value mat.Matrix
		values []complex128
		isPsd bool
		factorizeOk bool
	}{
		{values: []complex128{0,5}, isPsd: true, factorizeOk: true, value: dummyMatrix},
		{values: []complex128{0,-5}, isPsd: false, factorizeOk: true, value: dummyMatrix},
		{values: []complex128{0,complex(5, 7)}, isPsd: false, factorizeOk: true, value: dummyMatrix},
		{values: []complex128{}, isPsd: false, factorizeOk: true, value: mat.NewDense(2, 2, []float64{0,1,0,0})},
	}

	for _, table := range tables {
	candidate := PositiveSemidefiniteCandidate{ Value: table.value }
    isPsd, err := candidate.IsPositiveSemidefinite(createEigenMock(table.factorizeOk, table.values))
        if err != nil {
			t.Errorf("IsPositiveSemidefinite returned unexpected error: %v", err)
		}
		if isPsd != table.isPsd {
			t.Errorf("IsPositiveSemidefinite was incorrect, got: %t, want: %t.", isPsd, table.isPsd)
		}
	}
}

func TestPsdFactorizeNotOk(t *testing.T) {
	candidate := PositiveSemidefiniteCandidate{ Value: dummyMatrix }
	isPsd, err := candidate.IsPositiveSemidefinite(createEigenMock(false, []complex128{}))
	expectedError := fmt.Errorf("eigen: Factorize unsuccessful %v", mat.Formatted(candidate.Value, mat.Prefix("    "), mat.Squeeze()))
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Wrong error returned, got: %t, want: %t.", err, expectedError)
	}
	if isPsd {
		t.Errorf("IsPositiveSemidefinite was incorrect, got: %t, want: %t.", isPsd, false)
	}
}