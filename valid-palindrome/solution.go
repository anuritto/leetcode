package validpalindrome

import "unicode"

func isPalindrome(s string) bool {
	rs := []rune{}
	for _, l := range s {
		if unicode.IsLetter(l) || unicode.IsNumber(l) {
			rs = append(rs, unicode.ToLower(l))
		}
	}
	lowerStr := string(rs)

	stringLengh := len(lowerStr)
	for i := 0; i < stringLengh/2; i++ {
		if lowerStr[i] != lowerStr[stringLengh - 1 -i] {
			return false
		}
	}
	return true
}
