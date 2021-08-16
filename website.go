package faker

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

func BuildWebsite(domain string) string {
	data := []string{
		"https://",
		"https://www.",
		"http://",
		"http://www.",
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))] + domain
}