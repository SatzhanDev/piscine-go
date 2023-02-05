package piscine

func SplitWhiteSpaces(s string) []string {
	index := 0
	count := 0
	word := ""
	for i, v := range s {
		if v == ' ' && s[i+1] != ' ' {
			count++
		}
	}
	result := make([]string, count+1)
	for _, r := range s {
		if separator(r) {
			if word != "" {
				result[index] = word
				index++
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	size := 0
	for z := range result {
		size++
		z++
	}
	if word != "" {
		result[size-1] = word
	}
	return result
}

func separator(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
