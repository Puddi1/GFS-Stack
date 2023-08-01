package utils

import "net/mail"

// ValidMailAddress will check the validity of an email address
func ValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}
