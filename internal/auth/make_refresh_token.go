package auth

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func MakeRefreshToken() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("rand.Read gave an error: %s", err)
	}
	encodedStr := hex.EncodeToString(key)

	return encodedStr
}
