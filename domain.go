package faker

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

func PickUpDomain() string {
	data := []string{
		"example.com",
		"example.org",
		"example.net",
		"example.edu",
		"example.jp",
		"example.co.jp",
		"example.ne.jp",
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))]
}
