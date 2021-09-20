package website

import (
	"math/rand"

	"github.com/walkersumida/jaker/helpers/seed"
)

func Gen(domain string) string {
	data := []string{
		"https://",
		"https://www.",
		"http://",
		"http://www.",
	}
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))] + domain
}
