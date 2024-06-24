package utils

import "regexp"

// Regular expression for validating an email
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

// ValidateEmail checks if the provided email is in a valid format
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}
