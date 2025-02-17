package utility

import (
	"crypto/rand"
)

func GenerateVerificationCode(length int) string {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		Clog(err)
	}

	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b)
}
