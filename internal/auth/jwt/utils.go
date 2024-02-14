package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func generateSecretKey(length int) (string, error) {

	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	secretKey := base64.URLEncoding.EncodeToString(key)

	return secretKey, nil
}
