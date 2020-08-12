package emailproviders

import (
	"fmt"
	"strings"
)

var empty = struct{}{}

func Domain(email string) string {
	idx := strings.LastIndexByte(email, '@')
	if idx == -1 {
		panic(fmt.Errorf("email %s is missing domain", email))
	}
	return email[idx+1:]
}

func IsPublicEmail(email string) bool {
	_, found := publicEmailServices[Domain(email)]
	return found
}

func IsDisposableEmail(email string) bool {
	_, found := disposableEmailServices[Domain(email)]
	return found
}

func IsWorkEmail(email string) bool {
	domain := Domain(email)

	if _, found := publicEmailServices[domain]; found {
		return false
	}
	if _, found := disposableEmailServices[domain]; found {
		return false
	}
	return true
}
