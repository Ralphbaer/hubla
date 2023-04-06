package common

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateSHA256Hash(binaryData []byte) string {
	hasher := sha256.New()
	hasher.Write(binaryData)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
