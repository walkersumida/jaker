package faker

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

func PickUpFirstName() NameStruct {
	data := []NameStruct{
		{ JaKanji: "太郎", JaHira: "たろう", En: "taro" },
		{ JaKanji: "一郎", JaHira: "いちろう", En: "ichiro" },
		{ JaKanji: "次郎", JaHira: "じろう", En: "jiro" },
		{ JaKanji: "三郎", JaHira: "さぶろう", En: "saburo" },
		{ JaKanji: "四郎", JaHira: "しろう", En: "shiro" },
		{ JaKanji: "五郎", JaHira: "ごろう", En: "goro" },
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	return data[rand.Intn(len(data))]
}