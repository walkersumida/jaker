package jaker

import (
	guuid "github.com/google/uuid"
	"github.com/walkersumida/jaker/gen/company"
	"github.com/walkersumida/jaker/gen/domain"
	"github.com/walkersumida/jaker/gen/name"
	"github.com/walkersumida/jaker/gen/text"
	"github.com/walkersumida/jaker/gen/website"
	"github.com/walkersumida/jaker/helpers/upper"
)

type EmailStruct struct {
	Email          string
	EmailLocalPart string
	EmailDomain    string
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

func profile() ProfileStruct {
	var p ProfileStruct
	domain := domain.Gen()
	firstName := name.FirstGen()
	lastName := name.LastGen()
	email := BuildEmail(firstName.En, lastName.En, domain)
	company := company.Build(lastName)

	p.EnFirstName = firstName.En
	p.JaHiraFirstName = firstName.JaHira
	p.JaKanjiFirstName = firstName.JaKanji
	p.EnLastName = lastName.En
	p.JaHiraLastName = lastName.JaHira
	p.JaKanjiLastName = lastName.JaKanji
	p.EnFullName = upper.First(firstName.En) + " " + upper.First(lastName.En)
	p.JaHiraFullName = lastName.JaHira + " " + firstName.JaHira
	p.JaKanjiFullName = lastName.JaKanji + " " + firstName.JaKanji
	p.EnCompany = company.En
	p.JaCompany = company.Ja
	p.Website = website.Gen(domain)
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
