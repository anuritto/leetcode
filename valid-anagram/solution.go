package validanagram

func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }

	sm:= stringToMap(s)
	tm:= stringToMap(t)

	for l := range(sm) {
		if sm[l] != tm[l] {
			return false
		}
	}

	return true
}

func stringToMap(s string) map[string]int {
	m := map[string]int{}
	for i := 0; i < len(s); i++ {
		l:= string(s[i])
		m[l]++
	}
	return m
}