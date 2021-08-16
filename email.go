package faker

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

func BuildEmail(firstName string, lastName string, domain string) EmailStruct {
	localPart := firstName + "." + lastName
	email := localPart + "@" + domain

	data := []EmailStruct{
		{ Email: email, EmailLocalPart: localPart, EmailDomain: domain },
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))]
}