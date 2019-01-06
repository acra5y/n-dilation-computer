package ndilation

import (
	"github.com/acra5y/n-dilation-computer/positiveDefinite"
	"gonum.org/v1/gonum/mat"
)

type Dilation struct {
	N int32
	nDilation *mat.Dense
}

type isPositiveDefinite func(positiveDefinite.EigenComputer) (bool, error)

type squareRoot func(*mat.Dense) (*mat.Dense, error)

type newBlockMatrixFromSquares func([][]*mat.Dense) (*mat.Dense, error)

func (Dilation *Dilation) Value() *mat.Dense {
	return Dilation.nDilation
}

func (Dilation *Dilation) UnitaryNDilation(isPD isPositiveDefinite, sqrt squareRoot, newBlockMatrix newBlockMatrixFromSquares, t *mat.Dense) (error) {
	dummy := mat.NewDense(2, 2, nil)
	Dilation.nDilation = dummy
	return nil
}