# fastrand [![Circle CI](https://circleci.com/gh/segmentio/fastrand.svg?style=shield&circle-token=8c82ef91e2078833770b28936dea96fb22246e10)](https://circleci.com/gh/segmentio/fastrand) [![Go Report Card](https://goreportcard.com/badge/github.com/segmentio/fastrand)](https://goreportcard.com/report/github.com/segmentio/fastrand) [![GoDoc](https://godoc.org/github.com/segmentio/fastrand?status.svg)](https://godoc.org/github.com/segmentio/fastrand)

This package provides the implementation of a pseudo-random number source which
can be used with the standard `rand.Rand` type, and supports re-seeding with no
overhead.

Note that the pseudo-random number source in this package is not safe to use for
security-sensitive work.

## Motivation

Deterministic random number generation is a powerful tool, we use it to build
systems that have testable random behaviors. While the variety of outputs they
produce can be seen as randomly distributed, a given set of inputs will always
produce the same outputs.

The default pseudo-random number source in the `math/rand` package supports
building such systems because the sequence of numbers it generates is determined
by the value it was seeded with. However, seeding the source is an **extremely
expensive** operation, taking many microseconds of compute time, which doesn't
make it a good fit when the program needs to re-seed the PRNG based on inputs it
receives thousands of times per second, very quickly most of the CPU time is
spent re-seeding the PRNG.

This package provides an alternative pseudo-random number source which may not
provide a distribution that is as uniform as the one in the standard library,
but works well in practice, and supports re-seeding with zero overhead on CPU
time.

## Usage

```go
package main

import (
    "fmt"
    "math/rand"

    "github.com/segmentio/fastrand"
)

func main() {
    s := fastrand.NewSource(0)
    r := rand.New(s)
    fmt.Println(r.Int())
}
```
