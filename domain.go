package faker

import (
	"math/rand"

	"github.com/walkersumida/faker/helpers/seed"
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
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))]
}
