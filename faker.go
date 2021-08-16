package faker

import "strings"

type NameStruct struct {
	En string
	JaHira string
	JaKanji string
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
}

func firstName() NameStruct {
	return PickUpFirstName()
}

func lastName() NameStruct {
	return PickUpLastName()
}

func upper(str string) string {
	top := str[:1]
	others := str[1:]
	return strings.ToUpper(top) + others
}

func profile() ProfileStruct {
	var p ProfileStruct
	firstName := PickUpFirstName()
	lastName := PickUpLastName()
	p.EnFirstName = firstName.En
	p.JaHiraFirstName = firstName.JaHira
	p.JaKanjiFirstName = firstName.JaKanji
	p.EnLastName = lastName.En
	p.JaHiraLastName = lastName.JaHira
	p.JaKanjiLastName = lastName.JaKanji
	p.EnFullName = upper(firstName.En) + " " + upper(lastName.En)
	p.JaHiraFullName = lastName.JaHira + " " + firstName.JaHira
	p.JaKanjiFullName = lastName.JaKanji + " " + firstName.JaKanji

	return p
}

var Profile = profile()
