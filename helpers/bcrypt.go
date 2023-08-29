package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	salt := 8
	passwordBytes := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(passwordBytes, salt)

	return string(hash)
}

func ComparePassword(hashedPassword, password []byte) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)

	return err == nil
}
