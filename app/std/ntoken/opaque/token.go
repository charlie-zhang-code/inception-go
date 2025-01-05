package opaque

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// HashToken returns a hash of the token using the secret.
func HashToken(token string, secret []byte) (string, error) {
	h := hmac.New(sha256.New, secret)

	h.Write([]byte(token))

	return hex.EncodeToString(h.Sum(nil)), nil
}

// CheckHashTokenSecret checks if the hash of the token matches the secret.
func CheckHashTokenSecret(hashedToken string, rawToken string, secret []byte) bool {
	// Hash the raw token with the same secret to get the expected hash
	expectedHash, err := HashToken(rawToken, secret)
	if err != nil {
		return false // or handle the error as appropriate for your application
	}

	// Compare the hashes in constant time to prevent timing attacks
	return hmac.Equal([]byte(hashedToken), []byte(expectedHash))
}

// GenerateRandomToken generates a random token of the given length.
func GenerateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
