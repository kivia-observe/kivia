package apikey

import (
	"crypto/rand"
	"encoding/base64"
)

func generateAPIKey() (string, error) {
	bytes := make([]byte, 32) // 256-bit key
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	key := base64.RawURLEncoding.EncodeToString(bytes)

	return "kivia_" + key, nil
}
