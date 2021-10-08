package utils

import "net/mail"

func IsValidEmail(emailString string) bool {
	_, err := mail.ParseAddress(emailString)
	return err == nil
}
