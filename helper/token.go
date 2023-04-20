// Package helper contains some random functions
package helper

import (
	"crypto/rand"
	"encoding/hex"
)

// RandToken generates a random hex value.
func RandToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
