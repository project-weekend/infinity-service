package user

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// generateSecureToken creates a cryptographically secure token using HMAC-SHA256
func generateSecureToken(secret string) (string, string, error) {
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	rawToken := base64.URLEncoding.EncodeToString(randomBytes)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(rawToken))
	hashedToken := hex.EncodeToString(h.Sum(nil))

	return rawToken, hashedToken, nil
}

func hashToken(rawToken, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(rawToken))
	return hex.EncodeToString(h.Sum(nil))
}
