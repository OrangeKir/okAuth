package controllers

import (
	"crypto/sha512"
	"encoding/base64"
)

func HashPassword(password string) string {
	passwordBytes := []byte(password)

	sha512Hasher := sha512.New()
	sha512Hasher.Write(passwordBytes)

	hashedPasswordBytes := sha512Hasher.Sum(nil)
	base64EncodedPasswordHash := base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}
