package positiveSemidefinite

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
    "math/cmplx"
)

type PositiveSemidefiniteCandidate struct {
    Value *mat.Dense
}

type EigenComputer interface {
    Factorize(a mat.Matrix, left, right bool) bool
    Values(dst []complex128) []complex128
}

func isSymmetric(a mat.Matrix) bool {
    return mat.Equal(a, a.T())
}

func (candidate PositiveSemidefiniteCandidate) IsPositiveSemidefinite(eigen EigenComputer) (isPositiveSemidefinite bool, err error) {
    m, n := candidate.Value.Dims()
    c := mat.NewDense(m, n, nil)
    c.Clone(candidate.Value)
    err = nil
    isPositiveSemidefinite = true
    if !isSymmetric(candidate.Value) {
        isPositiveSemidefinite = false
        return
    }
    ok := eigen.Factorize(c, false, false)

    if ok {
        for _, val := range eigen.Values(nil) {
            _, theta := cmplx.Polar(val)
            if theta != 0 {
                isPositiveSemidefinite = false
                return
            }
        }
        return
    }
    return false, fmt.Errorf("eigen: Factorize unsuccessful %v", mat.Formatted(candidate.Value, mat.Prefix("    "), mat.Squeeze()))
}
