package faker

import (
	"math/rand"

	"github.com/walkersumida/faker/helpers/seed"
)

func BuildWebsite(domain string) string {
	data := []string{
		"https://",
		"https://www.",
		"http://",
		"http://www.",
	}
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))] + domain
}