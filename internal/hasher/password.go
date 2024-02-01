package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

var salt = []byte("Che_problema???_Ya_problema!!!")

func Encrypt(password string) (string, error) {
	crypted, err := bcrypt.GenerateFromPassword(append([]byte(password), salt...), 3)
	return string(crypted), err
}

func CorrectPassword(crypted string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(crypted), []byte(password))
	return err == nil
}
