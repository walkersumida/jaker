package seed

import (
	crand "crypto/rand"
	"math"
	"math/big"
)

func Int64() int64 {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))

	return seed.Int64()
}
