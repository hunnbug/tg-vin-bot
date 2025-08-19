package models

import (
	"regexp"
	"strings"
)

func IsVIN(VIN string) bool {
	if len(VIN) != 17 {
		return false
	}

	validChars := "0123456789ABCDEFGHJKLMNPRSTUVWXYZ"
	Upper := strings.ToUpper(VIN)
	for _, c := range Upper {
		if !strings.ContainsRune(validChars, c) {
			return false
		}
	}

	vinRegex := regexp.MustCompile(`^[A-HJ-NPR-Z0-9]{17}$`)
	return vinRegex.MatchString(Upper)
}
