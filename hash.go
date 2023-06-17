package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashingSHA256 fungsi untuk melakukan hash SHA256
func hashingSHA256(data string) string {
	byteData := []byte(data)
	sha256Hash := sha256.New()
	sha256Hash.Write(byteData)
	hashResult := sha256Hash.Sum(nil)
	hashString := hex.EncodeToString(hashResult)
	return hashString
}
