package main

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"fmt"
)

func printDivider() {
	fmt.Printf("%s\n", "--------------------------------")
}

func printM(a mat.Matrix) {
	fa := mat.Formatted(a, mat.Prefix("    "), mat.Squeeze())

	fmt.Printf("\na = %v\n\n", fa)
	printDivider()
}

func main() {
	dimension := 6
	data := make([]float64, 36)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	a := mat.NewDense(dimension, dimension, data)

	tr := mat.Trace(a)

	fmt.Printf("%g\n", tr)
	printDivider()
	printM(a)
	printDivider()

	values := mat.Eigen{}
	values.Factorize(a, false, true)

	printDivider()
	for _, v := range values.Values(nil) {
		fmt.Printf("%g\n", v)
	}

	var eigen mat.Eigen
	eigen.Factorize(a, false, false)

	for _, v := range eigen.Values(nil) {
		fmt.Printf("%g\n", v)
	}
}
