package squareRoot

import (
    "gonum.org/v1/gonum/mat"
    "testing"
)

var dummyMatrix = mat.NewDense(2, 2, nil)

func TestCalculate(t *testing.T) {
    tables := []struct {
        desc string
        value mat.Matrix
    }{
        {value: dummyMatrix, desc: "template"},
    }

    for _, table := range tables {

        t.Run(table.desc, func(t *testing.T) {
            t.Parallel()
            squareRoot := SquareRoot{
                Value: table.value,
            }
            _, err := squareRoot.Calculate()

            if err == nil {
                t.Errorf("No error thrown.")
            }
        })
    }
}
