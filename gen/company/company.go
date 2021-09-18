package company

import (
	"math/rand"

	"github.com/walkersumida/jaker/gen/name"
	"github.com/walkersumida/jaker/helpers/seed"
	"github.com/walkersumida/jaker/helpers/upper"
)

type CompanyStruct struct {
	En string
	Ja string
}

func Build(name name.NameStruct) CompanyStruct {
	data := []CompanyStruct{
		{En: "Capital Partners", Ja: "キャピタルパートナーズ"},
		{En: "Technologies", Ja: "テクノロジーズ"},
		{En: "Systems", Ja: "システムズ"},
		{En: "Companies", Ja: "カンパニーズ"},
		{En: "Holdings", Ja: "ホールディングス"},
		{En: "Entertainment", Ja: "エンターテインメント"},
		{En: "Financial Group", Ja: "フィナンシャルグループ"},
		{En: "Group", Ja: "グループ"},
		{En: "Communications", Ja: "コミュニケーションズ"},
		{En: "Foods", Ja: "食品"},
		{En: "Real Estate Development", Ja: "不動産"},
		{En: "Trading Company", Ja: "商事"},
	}
	rand.Seed(seed.Int64())

	selectedCompany := data[rand.Intn(len(data))]
	selectedCompany.En = upper.First(name.En) + " " + selectedCompany.En + ", Inc."
	selectedCompany.Ja = "株式会社" + name.JaKanji + selectedCompany.Ja

	return selectedCompany
}
