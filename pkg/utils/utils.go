package utils

import (
	"regexp"
	"strings"
)

func ValidaCPF(cpf string) bool {
	// remove non-numeric characters
	cpf = strings.Join(regexp.MustCompile("[0-9]+").FindAllString(cpf, -1), "")
	if len(cpf) != 11 {
		return false
	}

	// simple cpf validation
	var sum int
	for i := 0; i < 9; i++ {
		digit := int(cpf[i] - '0')
		sum += digit * (10 - i)
	}
	firstDigit := (sum * 10 % 11) % 10

	sum = 0
	for i := 0; i < 10; i++ {
		digit := int(cpf[i] - '0')
		sum += digit * (11 - i)
	}
	secondDigit := (sum * 10 % 11) % 10

	return firstDigit == int(cpf[9]-'0') && secondDigit == int(cpf[10]-'0')
}
