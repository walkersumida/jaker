package faker

import "strings"

type NameStruct struct {
	En string
	JaHira string
	JaKanji string
}

type EmailStruct struct {
	Email string
	EmailLocalPart string
	EmailDomain string
}

type ProfileStruct struct {
	EnFirstName string
	EnLastName string
	JaHiraFirstName string
	JaHiraLastName string
	JaKanjiFirstName string
	JaKanjiLastName string
	EnFullName string
	JaHiraFullName string
	JaKanjiFullName string
	Website string
	EmailStruct
}

func upper(str string) string {
	top := str[:1]
	others := str[1:]
	return strings.ToUpper(top) + others
}

func profile() ProfileStruct {
	var p ProfileStruct
	domain := PickUpDomain()
	firstName := PickUpFirstName()
	lastName := PickUpLastName()
	email := BuildEmail(firstName.En, lastName.En, domain)

	p.EnFirstName = firstName.En
	p.JaHiraFirstName = firstName.JaHira
	p.JaKanjiFirstName = firstName.JaKanji
	p.EnLastName = lastName.En
	p.JaHiraLastName = lastName.JaHira
	p.JaKanjiLastName = lastName.JaKanji
	p.EnFullName = upper(firstName.En) + " " + upper(lastName.En)
	p.JaHiraFullName = lastName.JaHira + " " + firstName.JaHira
	p.JaKanjiFullName = lastName.JaKanji + " " + firstName.JaKanji
	p.Website = BuildWebsite(domain)
	p.Email = email.Email
	p.EmailLocalPart = email.EmailLocalPart
	p.EmailDomain = email.EmailDomain

	return p
}

var Profile = profile()
