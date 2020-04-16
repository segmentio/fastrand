// Package fastrand provides a port of the fnvrand algorithm found in
// https://stackoverflow.com/questions/664014/what-integer-hash-function-are-good-that-accepts-an-integer-hash-key/12996028#12996028
package fastrand

import (
	"math"
	"math/rand"
)

// NewSource constructs a new fnvrand source seeded with the value passed as
// argument.
func NewSource(seed int64) rand.Source64 {
	return &source{seed: uint64(seed)}
}

type source struct {
	seed uint64
}

func (rng *source) Seed(seed int64) {
	rng.seed = uint64(seed)
}

func (rng *source) Int63() int64 {
	// Ported from math/rand/rng.go
	const mask = math.MaxInt64 - 1
	return int64(rng.Uint64() & mask)
}

func (rng *source) Uint64() uint64 {
	hash := rng.seed
	rng.seed++

	hash = (hash ^ (hash >> 30)) * 0xbf58476d1ce4e5b9
	hash = (hash ^ (hash >> 27)) * 0x94d049bb133111eb
	hash = hash ^ (hash >> 31)

	return hash
}
