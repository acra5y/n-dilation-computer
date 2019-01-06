package dilation

import (
	"fmt"
	"github.com/acra5y/n-dilation-computer/eye"
	"github.com/acra5y/n-dilation-computer/positiveDefinite"
	"gonum.org/v1/gonum/mat"
)

type isPositiveDefinite func(positiveDefinite.EigenComputer, *mat.Dense) (bool, error)

type squareRoot func(*mat.Dense) (*mat.Dense, error)

type newBlockMatrixFromSquares func([][]*mat.Dense) (*mat.Dense, error)

func op (isPD isPositiveDefinite, sqrt squareRoot, t mat.Matrix) (*mat.Dense, error) {
	n, _ := t.Dims()
	eye := eye.OfDimension(n)

	t_times_t_transpose := mat.NewDense(n, n, nil)

	t_times_t_transpose.Product(t, t.T())

	opSquared := mat.NewDense(n, n, nil)

	opSquared.Sub(eye, t_times_t_transpose)

	if pd, _ := isPD(&mat.Eigen{}, opSquared); !pd {
		return mat.NewDense(0, 0, nil), fmt.Errorf("Input is not a contraction")
	}

	op, _ := sqrt(opSquared)

	return op, nil
}

func UnitaryNDilation(isPD isPositiveDefinite, sqrt squareRoot, newBlockMatrix newBlockMatrixFromSquares, t *mat.Dense) (*mat.Dense, error) {
	m, n := t.Dims()

	if m != n {
		return mat.NewDense(0,0, nil), fmt.Errorf("Matrix does not have square dimension")
	}

	d_t, err := op(isPD, sqrt, t)

	if err != nil {
		return mat.NewDense(0,0, nil), err
	}

	d_t_transpose, err := op(isPD, sqrt, t.T())

	if err != nil {
		return mat.NewDense(0,0, nil), err
	}

	unitary, err := newBlockMatrix([][]*mat.Dense{[]*mat.Dense{t, d_t,},[]*mat.Dense{d_t_transpose, mat.NewDense(m, m, nil),},})

	if err != nil {
		return mat.NewDense(0,0, nil), err
	}

    return unitary, nil
}