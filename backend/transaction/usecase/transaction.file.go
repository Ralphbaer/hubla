package usecase

import (
	"crypto/sha256"
	"encoding/hex"
)

func calculateSHA256Hash(binaryData []byte) string {
	hasher := sha256.New()
	hasher.Write(binaryData)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
