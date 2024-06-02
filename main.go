package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

func main() {
	const size = 4000 // Adjust this for larger matrices

	// Print the BLAS implementation being used
	fmt.Printf("In main, Using BLAS implementation: %T\n", blas64.Implementation())

	// Create random matrices
	dataA := make([]float64, size*size)
	dataB := make([]float64, size*size)
	for i := range dataA {
		dataA[i] = rand.Float64()
		dataB[i] = rand.Float64()
	}
	A := mat.NewDense(size, size, dataA)
	B := mat.NewDense(size, size, dataB)

	// Time the multiplication
	start := time.Now()
	var result mat.Dense
	result.Mul(A, B)
	elapsed := time.Since(start)
	fmt.Println("Matrix multiplication time:", elapsed)
}
