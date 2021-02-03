package mngapi

import (
	"math/rand"
	"time"
)

// AlphaNum contain ascii alpha and numeric chars
const AlphaNum = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomStringWithCharset returns a random string of given lenght from the given charset
func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString returns a random string of given lenght
func RandomString(length int) string {
	return RandomStringWithCharset(length, AlphaNum)
}
