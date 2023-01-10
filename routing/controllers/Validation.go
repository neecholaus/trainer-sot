package controllers

import (
	"fmt"
	"regexp"
)

type Validator struct{}

func (_ *Validator) IsEmail(toValidate string) bool {
	if toValidate == "" {
		return false
	}

	match, err := regexp.Match(`^\w+@\w{2,}.\w{2}`, []byte(toValidate))
	if err != nil {
		fmt.Printf("error while validating as email: %s\n", err)
		return false
	}

	return match
}
