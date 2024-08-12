package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	// Seed the random number generator using the current time
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandOwner() int64 {
	return RandomInt(1, 4)
}

func RandBalance() int64 {
	return RandomInt(10, 1000)
}

func RandCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "GBP", "JPY", "KES"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandUserName() string {
	return RandomString(6)
}

func RandUserEmail() string {
	return RandomString(6) + "@gmail.com"
}
