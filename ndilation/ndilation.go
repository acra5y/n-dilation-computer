package ndilation

import (
	"github.com/acra5y/n-dilation-computer/positiveDefinite"
	"gonum.org/v1/gonum/mat"
)

type Dilation struct {
	N int32
	nDilation *mat.Dense
}

type validation interface {
	IsPositiveDefinite(positiveDefinite.EigenComputer) (bool, error)
}

type squareRoot interface {
	Calculate(*mat.Dense) (*mat.Dense, error)
}

type blockMatrix interface {
	NewBlockMatrixFromSquares([][]*mat.Dense) (*mat.Dense, error)
}

func (Dilation *Dilation) Value() *mat.Dense {
	return Dilation.nDilation
}

func (Dilation *Dilation) UnitaryNDilation(validation validation, squareRoot squareRoot, blockMatrix blockMatrix, t *mat.Dense) (error) {
	dummy := mat.NewDense(2, 2, nil)
	Dilation.nDilation = dummy
	return nil
}