package piscine

func isWhiteSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	wordStart := 0
	inWord := false

	for i := 0; i < len(s); i++ {
		if isWhiteSpace(s[i]) {
			if inWord {
				result = append(result, s[wordStart:i])
				inWord = false
			}
		} else {
			if !inWord {
				wordStart = i
				inWord = true
			}
		}
	}

	if inWord {
		result = append(result, s[wordStart:])
	}

	return result
}
