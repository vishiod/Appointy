package utils

import (
	"crypto/hmac"
	"crypto/sha256" // Irrevocable
	"encoding/hex"
	"hash"
)

var secretKeyPerApplication = "o$cfyg.C5Tg5a05Llyv9qdC5&JGtQH&$"

func ComputeHmac256(message string) hash.Hash {

	key := []byte(secretKeyPerApplication)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return h
}

func HMACToString(messageHash hash.Hash) string {
	return hex.EncodeToString(messageHash.Sum(nil))
}

func IsHMACEqual(message1 hash.Hash, message2 hash.Hash) bool {
	return hmac.Equal(message1.Sum(nil), message2.Sum(nil))
}
