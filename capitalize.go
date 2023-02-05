package piscine

func Capitalize(s string) string {
	a := []rune(s)
	for index, letter := range a {
		if check(letter) {
			if index == 0 || check(a[index-1]) == false {
				if letter >= 'a' && letter <= 'z' {
					a[index] = letter - 32
				}
			} else {
				if letter >= 'A' && letter <= 'Z' {
					a[index] = letter + 32
				}
			}
		}
	}
	return string(a)
}

func check(r rune) bool {
	if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
