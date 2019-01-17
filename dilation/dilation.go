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

func defectOperator (isPD isPositiveDefinite, sqrt squareRoot, t mat.Matrix) (*mat.Dense, error) {
    n, _ := t.Dims()
    eye := eye.OfDimension(n)

    tTimesTTransposed := mat.NewDense(n, n, nil)

    tTimesTTransposed.Product(t, t.T())

    defectSquared := mat.NewDense(n, n, nil)

    defectSquared.Sub(eye, tTimesTTransposed)

    if pd, _ := isPD(&mat.Eigen{}, defectSquared); !pd {
        return mat.NewDense(0, 0, nil), fmt.Errorf("Input is not a contraction")
    }

    defectOp, _ := sqrt(defectSquared)

    return defectOp, nil
}

func negativeTranspose(t *mat.Dense) *mat.Dense {
    m, n := t.Dims()
    data := make([]float64, m * n)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            data[m * i + j] = (-1) * t.At(j, i)
        }
    }
    return mat.NewDense(m, n, data)
}

// See E. Levy und O. M. Shalit: Dilation theory in finite dimensions: the possible, the impossible and the unknown. Rocky Mountain J. Math., 44(1):203-221, 2014

func UnitaryNDilation(isPD isPositiveDefinite, sqrt squareRoot, newBlockMatrix newBlockMatrixFromSquares, t *mat.Dense) (*mat.Dense, error) {
    m, n := t.Dims()

    if m != n {
        return mat.NewDense(0,0, nil), fmt.Errorf("Matrix does not have square dimension")
    }

    defect, err := defectOperator(isPD, sqrt, t)

    if err != nil {
        return mat.NewDense(0,0, nil), err
    }

    defectOfTransposed, err := defectOperator(isPD, sqrt, t.T())

    if err != nil {
        return mat.NewDense(0,0, nil), err
    }

    unitary, err := newBlockMatrix([][]*mat.Dense{[]*mat.Dense{t, defectOfTransposed,},[]*mat.Dense{defect, negativeTranspose(t),},})

    if err != nil {
        return mat.NewDense(0,0, nil), err
    }

    return unitary, nil
}
