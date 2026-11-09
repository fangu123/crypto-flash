package util

import (
	"errors"
	fault "go-admin/common/errors"
	"regexp"
)

func ValidatePassword(pw, lang string) error {
	if len(pw) < 12 {
		return errors.New(fault.GetErrorMessage("password_length_at_least_12", lang))
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pw)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(pw)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(pw)
	hasSpecial := regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(pw)
	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return errors.New(fault.GetErrorMessage("password_must_contain_uppercase_lowercase_digit_special_character", lang))
	}
	return nil
}
