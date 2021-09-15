package jaker

import (
	"strings"

	guuid "github.com/google/uuid"
	"github.com/walkersumida/jaker/gen/text"
)

type NameStruct struct {
	En      string
	JaHira  string
	JaKanji string
}

type EmailStruct struct {
	Email          string
	EmailLocalPart string
	EmailDomain    string
}

type CompanyStruct struct {
	En string
	Ja string
}

type ProfileStruct struct {
	EnFirstName      string
	EnLastName       string
	JaHiraFirstName  string
	JaHiraLastName   string
	JaKanjiFirstName string
	JaKanjiLastName  string
	EnFullName       string
	JaHiraFullName   string
	JaKanjiFullName  string
	EnCompany        string
	JaCompany        string
	Website          string
	EmailStruct
}

func Upper(str string) string {
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
	company := BuildCompany(lastName)

	p.EnFirstName = firstName.En
	p.JaHiraFirstName = firstName.JaHira
	p.JaKanjiFirstName = firstName.JaKanji
	p.EnLastName = lastName.En
	p.JaHiraLastName = lastName.JaHira
	p.JaKanjiLastName = lastName.JaKanji
	p.EnFullName = Upper(firstName.En) + " " + Upper(lastName.En)
	p.JaHiraFullName = lastName.JaHira + " " + firstName.JaHira
	p.JaKanjiFullName = lastName.JaKanji + " " + firstName.JaKanji
	p.EnCompany = company.En
	p.JaCompany = company.Ja
	p.Website = BuildWebsite(domain)
	p.Email = email.Email
	p.EmailLocalPart = email.EmailLocalPart
	p.EmailDomain = email.EmailDomain

	return p
}

func uuid() string {
	uuidObj, _ := guuid.NewRandom()
	uuid := uuidObj.String()

	return uuid
}

func Text(base string, length int) string {
	txt := text.New()
	txt.Base = base
	txt.Size = length

	return txt.Gen()
}

var Profile = profile()
var Uuid = uuid()
