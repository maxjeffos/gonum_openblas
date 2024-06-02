//go:build !openblas

package main

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/blas/gonum"
)

func init() {
	fmt.Println("init() setting BLAS implementation: gonum.Implementation")

	// Enable Go native BLAS implementation
	blas64.Use(gonum.Implementation{})
}
