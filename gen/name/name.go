package name

import (
	"math/rand"

	"github.com/walkersumida/jaker/helpers/seed"
)

type NameStruct struct {
	En      string
	JaHira  string
	JaKanji string
}

func FirstGen() NameStruct {
	data := []NameStruct{
		{JaKanji: "太郎", JaHira: "たろう", En: "taro"},
		{JaKanji: "一郎", JaHira: "いちろう", En: "ichiro"},
		{JaKanji: "次郎", JaHira: "じろう", En: "jiro"},
		{JaKanji: "三郎", JaHira: "さぶろう", En: "saburo"},
		{JaKanji: "四郎", JaHira: "しろう", En: "shiro"},
		{JaKanji: "五郎", JaHira: "ごろう", En: "goro"},
	}
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))]
}

func LastGen() NameStruct {
	data := []NameStruct{
		{JaKanji: "佐藤", JaHira: "さとう", En: "sato"},
		{JaKanji: "鈴木", JaHira: "すずき", En: "suzuki"},
		{JaKanji: "高橋", JaHira: "たかはし", En: "takahashi"},
		{JaKanji: "田中", JaHira: "たなか", En: "tanaka"},
		{JaKanji: "伊藤", JaHira: "いとう", En: "ito"},
		{JaKanji: "渡辺", JaHira: "わたなべ", En: "watanabe"},
		{JaKanji: "山本", JaHira: "やまもと", En: "yamamoto"},
		{JaKanji: "中村", JaHira: "なかむら", En: "nakamura"},
		{JaKanji: "小林", JaHira: "こばやし", En: "kobayashi"},
		{JaKanji: "加藤", JaHira: "かとう", En: "kato"},
	}
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))]
}
