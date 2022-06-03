package aes

import (
	"crypto/sha256"
	//"fmt"
	"encoding/hex"
)

func Sha256t(t string) string {
	sum := sha256.Sum256([]byte(t))
	return hex.EncodeToString(sum[:])
}