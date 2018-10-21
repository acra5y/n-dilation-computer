package positiveSemidefiniteTester

import (
	"gonum.org/v1/gonum/mat"
	"math/cmplx"
)

type PositiveSemidefiniteCandidate struct {
	Value mat.Matrix
}

type EigenComputer interface {
	Factorize(a mat.Matrix, left, right bool) bool
	Values(dst []complex128) []complex128
}

func isSymmetric(a mat.Matrix) bool {
	return mat.Equal(a, a.T())
}

func (candidate PositiveSemidefiniteCandidate) IsPositiveSemidefinite(eigen EigenComputer) (isPositiveSemidefinite bool) {
	if !isSymmetric(candidate.Value) {
		return false
	}
	eigen.Factorize(candidate.Value, false, false)
	isPositiveSemidefinite = true

	for _, val := range eigen.Values(nil) {
		_, theta := cmplx.Polar(val)
		if theta != 0 {
			isPositiveSemidefinite = false
			return
		}
	}
	return
}
