package validpalindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for !(unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i]))) {
			i++
		}

		for !(unicode.IsLetter(rune(s[j])) || unicode.IsNumber(rune(s[j]))) {
			j--
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
	}

	return true
}
