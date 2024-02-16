package validator

import (
	"braincome/internal/models"
	"net/mail"
	"regexp"
	"unicode"
	"unicode/utf8"
)

const (
	MsgNumberExists       = "phone number already exists"
	MsgEmailExists        = "email already exists"
	MsgNameExists         = "name already exists"
	MsgInvalidEmail       = "write correct email"
	MsgInvalidFirstName   = "name must consist only of English letters."
	MsgInvalidLastName    = "name must consist only of English letters."
	MsgInvalidPass        = "password must contain letters, numbers and must be at least 6 characters."
	MsgUserNotFound       = "user not found"
	MsgPassDontMatch      = "the passwords don't match"
	MsgNotCorrectPassword = "password is not correct"
)

func GetErrMsgs(m models.User) map[string]string {
	errmsgs := make(map[string]string)
	if !isValidEmail(m.Email) {
		errmsgs["email"] = MsgInvalidEmail
	}
	if !isValidName(m.First_name) {
		errmsgs["first_name"] = MsgInvalidFirstName
	}
	if !isValidName(m.Last_name) {
		errmsgs["last_name"] = MsgInvalidLastName
	}
	if !isValidPassword(m.Password) {
		errmsgs["pass"] = MsgInvalidPass
	}
	return errmsgs
}

func isValidEmail(email string) bool {
	rxEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidName(name string) bool {
	length := utf8.RuneCountInString(name)
	if length < 1 || length > 15 {
		return false
	}
	for _, char := range name {
		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true
}

func isValidPassword(pass string) bool {
	tests := []string{".{6,}", "[a-z]", "[0-9]"}
	for _, test := range tests {
		valid, _ := regexp.MatchString(test, pass)
		if !valid {
			return false
		}
	}
	return true
}
