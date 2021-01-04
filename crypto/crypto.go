package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

// SHA256 : SHA256 Hash string -> hexacode
func SHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	byteSlice := hash.Sum(nil)
	hex := fmt.Sprintf("%x", byteSlice)
	return hex
}

// MD5 : MD5 Hash string -> hexacode
func MD5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	byteSlice := hash.Sum(nil)
	hex := fmt.Sprintf("%x", byteSlice)
	return hex
}
