package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GetUnixTimestamp() int64 {
	return time.Now().UTC().UnixNano()
}

func GenerateUnixTimestamp() string {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	return timestamp
}

func GenerateNumeric(rangeNumber int) string {
	letterBytes := "0123456789"
	return generate(letterBytes, rangeNumber)
}

func GenerateAlphabet(rangeNumber int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return generate(letterBytes, rangeNumber)
}

func GenerateAlphaNumeric(rangeNumber int) string {
	letterBytes := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return generate(letterBytes, rangeNumber)
}

func GenerateUniqueCharacter(rangeNumber int) string {
	letterBytes := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`~!@#$%^&*()-_=+|]}[{';:/?.>,<"
	return generate(letterBytes, rangeNumber)
}

func generate(letterBytes string, rangeNumber int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, rangeNumber)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
