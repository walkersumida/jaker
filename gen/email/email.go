package email

import (
	"math/rand"

	"github.com/walkersumida/jaker/helpers/seed"
)

type EmailStruct struct {
	Email          string
	EmailLocalPart string
	EmailDomain    string
}

func Gen(firstName string, lastName string, domain string) EmailStruct {
	localPart := firstName + "." + lastName
	email := localPart + "@" + domain

	data := []EmailStruct{
		{Email: email, EmailLocalPart: localPart, EmailDomain: domain},
	}
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))]
}
