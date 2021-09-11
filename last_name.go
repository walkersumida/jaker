package jaker

import (
	"math/rand"

	"github.com/walkersumida/jaker/helpers/seed"
)

func PickUpLastName() NameStruct {
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
