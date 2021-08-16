package faker

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

func BuildCompany(name NameStruct) CompanyStruct {
	data := []CompanyStruct{
		{ En: "Capital Partners", Ja: "キャピタルパートナーズ" },
		{ En: "Technologies", Ja: "テクノロジーズ" },
		{ En: "Systems", Ja: "システムズ" },
		{ En: "Companies", Ja: "カンパニーズ" },
		{ En: "Holdings", Ja: "ホールディングス" },
		{ En: "Entertainment", Ja: "エンターテインメント" },
		{ En: "Financial Group", Ja: "フィナンシャルグループ" },
		{ En: "Group", Ja: "グループ" },
		{ En: "Communications", Ja: "コミュニケーションズ" },
		{ En: "Foods", Ja: "食品" },
		{ En: "Real Estate Development", Ja: "不動産" },
		{ En: "Trading Company", Ja: "商事" },
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	selectedCompany := data[rand.Intn(len(data))]
	selectedCompany.En = Upper(name.En) + " " + selectedCompany.En + ", Inc."
	selectedCompany.Ja = "株式会社" + name.JaKanji + selectedCompany.Ja

	return selectedCompany
}