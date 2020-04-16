package fastrand_test

import (
	"math/rand"
	"testing"

	"github.com/segmentio/fastrand"
)

func TestFastrand(t *testing.T) {
	const N = 1000000
	values := make(map[uint64]struct{}, N)
	source := fastrand.NewSource(0)

	for i := 0; i < N; i++ {
		x := source.Uint64()

		if _, exists := values[x]; exists {
			t.Errorf("%d was generated twice", x)
			break
		}

		values[x] = struct{}{}
	}
}

func BenchmarkNewSource(b *testing.B) {
	b.Run("math/rand", benchmarkMathRandNewSource)
	b.Run("fastrand", benchmarkFastrandNewSource)
}

func benchmarkMathRandNewSource(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rand.NewSource(1234567890)
	}
}

func benchmarkFastrandNewSource(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fastrand.NewSource(1234567890)
	}
}

func BenchmarkRandSource(b *testing.B) {
	b.Run("math/rand", benchmarkMathRandSource)
	b.Run("fastrand", benchmarkFastrandSource)
}

func benchmarkMathRandSource(b *testing.B) {
	benchmarkSource(b, rand.NewSource(0))
}

func benchmarkFastrandSource(b *testing.B) {
	benchmarkSource(b, fastrand.NewSource(0))
}

func benchmarkSource(b *testing.B, source rand.Source) {
	prng := rand.New(source)

	for i := 0; i < b.N; i++ {
		_ = prng.NormFloat64()
	}
}

func BenchmarkSeed(b *testing.B) {
	b.Run("math/rand", benchmarkMathRandSeed)
	b.Run("fastrand", benchmarkFastrandSeed)
}

func benchmarkMathRandSeed(b *testing.B) {
	r := rand.NewSource(0)

	for i := 0; i < b.N; i++ {
		r.Seed(1234567890)
	}
}

func benchmarkFastrandSeed(b *testing.B) {
	r := fastrand.NewSource(0)

	for i := 0; i < b.N; i++ {
		r.Seed(int64(i))
	}
}
