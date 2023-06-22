package utils

import (
	"math/rand"
)

var CodeNameList = []string{
	"apple",
	"banana",
	"carrot",
	"dragon",
	"elephant",
	"flower",
	"grape",
	"honey",
	"iguana",
	"jungle",
	"kiwi",
	"lemon",
	"mango",
	"nutmeg",
	"orange",
	"pineapple",
	"quinoa",
	"raspberry",
	"strawberry",
	"tulip",
	"umbrella",
	"violet",
	"watermelon",
	"xylophone",
	"yogurt",
	"zebra",
}

func GetCodeName() string {
	randomIDX := rand.Intn(len(CodeNameList))
	return CodeNameList[randomIDX]
}