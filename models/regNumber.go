package models

func IsRegNumber(number string) bool {
	if len(number) < 8 {
		return false
	} else {
		return true
	}
}
