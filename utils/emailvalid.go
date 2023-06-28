package utils

import "regexp"


func EmailValidate(email string) bool {
	EmailRegEx := "^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"
	if matched, _ := regexp.MatchString(EmailRegEx, email); !matched {
		return false
	}
	return true
} 

