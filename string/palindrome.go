package string

import "unicode/utf8"

func IsPalindrome(s string) bool {
	if !utf8.ValidString(s) {
		return false
	}

	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}
