package piscine

func AlphaCount(s string) int {
	count := 0
	for _, letter := range s {
		if (letter >= 97 && letter <= 122) || (letter >= 65 && letter <= 90) {
			count++
		}
	}
	return count
}
