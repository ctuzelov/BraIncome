package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

var salt = []byte("s3cr3t")

func Encrypt(password string) (string, error) {
	crypted, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	return string(crypted), err
}

func CorrectPassword(crypted string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(crypted), []byte(password))
	return err == nil
}
