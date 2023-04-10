package common

import (
	"crypto/sha256"
	"encoding/hex"
)

// CalculateSHA256Hash calculates the SHA-256 hash of the provided binary data and returns the result as a hexadecimal string.
func CalculateSHA256Hash(binaryData []byte) string {
	hasher := sha256.New()
	hasher.Write(binaryData)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
