# gonum_openblas

This repo is a starting point showing how to use [gonum](https://github.com/gonum/gonum) with OpenBLAS.

To run it, you can use:

```sh
# To run with the Go implementations:
go run .
# To run it with the OpenBLAS implementation:
go run -tags openblas .
```

Similarily to build, you can use:

```sh
# To run with the Go implementations:
go build .
# To run it with the OpenBLAS implementation:
go build -tags openblas .
```

The main() will do large matrix multiplication which shows the performance difference between using the Go vs OpenBLAS implementations.

## Cgo Config

In gonum_openblas.go, there's some Cgo configuration for using OpenBLAS. What I have is the simplest thing that is required to make this bootstrap work:

```go
/*
#cgo LDFLAGS: -L/opt/homebrew/opt/openblas/lib -lopenblas
*/
import "C"
```

Note that the path `/opt/homebrew/opt/openblas/lib` may vary depending on your system and how you installed OpenBLAS.

If you want to do more complex things, like calling openblas directly, you may include some additional configs, such as:

```go
/*
#cgo LDFLAGS: -L/opt/homebrew/opt/openblas/lib -lopenblas
#cgo CFLAGS: -I/opt/homebrew/opt/openblas/include
#include <cblas.h>
#include <openblas_config.h>
*/
import "C"
```
