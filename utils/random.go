package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// Hex : Random hexacode with length = numBytes * 2
func Hex(numBytes int) string {
	// new seed for random
	rand.Seed(time.Now().UnixNano())

	token := make([]byte, numBytes)
	rand.Read(token)

	// encode to hex with length = numBytes * 2
	return hex.EncodeToString(token)
}
