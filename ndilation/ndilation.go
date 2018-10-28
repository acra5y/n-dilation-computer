package ndilation

import (
	"gonum.org/v1/gonum/mat"
)

type Dilation struct {
	N int32
	nDilation *mat.Dense
}

func (Dilation *Dilation) Value() *mat.Dense {
	return Dilation.nDilation
}

func (Dilation *Dilation) unitaryNDilation(t *mat.Dense) (error) {
	dummy := mat.NewDense(2, 2, nil)
	Dilation.nDilation = dummy
	return nil
}