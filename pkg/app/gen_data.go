package app

import (
	"math/rand"
)

var letterRunes = []rune("0123456789")

func getRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func genCardsNumber() string {
	return getRandomString(16)
}

func genAccNumber() string {
	return "UA" + getRandomString(29)
}
