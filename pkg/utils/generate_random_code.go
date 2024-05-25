package utils

import (
	"math/rand"
	"time"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomCode(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	code := make([]rune, length)
	for i := range code {
		code[i] = letterRunes[randomGenerator.Intn(len(letterRunes))]
	}
	return string(code)
}
