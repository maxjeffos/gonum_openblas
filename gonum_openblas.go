//go:build openblas

package main

/*
#cgo LDFLAGS: -L/opt/homebrew/opt/openblas/lib -lopenblas
*/
import "C"
import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/netlib/blas/netlib"
)

func init() {
	fmt.Println("init() setting BLAS implementation: netlib.Implementation (OpenBLAS)")

	// Enable OpenBLAS by setting the BLAS implementation to netlib
	blas64.Use(netlib.Implementation{})
}
