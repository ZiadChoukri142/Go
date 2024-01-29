package piscine

func Verify(char rune) bool {
	if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	str := []rune(s)
	foundW := false
	for i, j := range str {
		if i == 0 && (j <= 'z' && j >= 'a') {
			str[i] = str[i] - 32
			foundW = true
			continue
		}
		if !Verify(j) {
			foundW = false
			continue
		}
		if Verify(j) && !foundW {
			if j <= 'z' && j >= 'a' {
				str[i] = str[i] - 32
			}
			foundW = true
			continue
		}
		if Verify(j) && (j <= 'Z' && j >= 'A') {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}
