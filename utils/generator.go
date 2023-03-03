package utils

import (
	"crypto/rand"
	"errors"
	"io"
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

func GenerateNumber(numberRange int) (*string, *error) {

	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, numberRange)

	n, err := io.ReadAtLeast(rand.Reader, b, numberRange)
	if err != nil {
		return nil, &err
	}

	if n != numberRange {
		newError := errors.New("invalid range")
		return nil, &newError
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	result := string(b)

	return &result, nil
}
