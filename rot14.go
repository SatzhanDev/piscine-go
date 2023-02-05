package piscine

func Rot14(str string) string {
	runes := []rune(str)
	var alphaLow [52]rune
	var aplhaUp [52]rune

	for i, j, l := 0, 'a', 26; j <= 'z'; i, j, l = i+1, j+1, l+1 {
		alphaLow[i] = j
		alphaLow[l] = j
	}

	for i, j, l := 0, 'A', 26; j <= 'Z'; i, j, l = i+1, j+1, l+1 {
		aplhaUp[i] = j
		aplhaUp[l] = j
	}

	for i, v := range runes {
		if v >= 'a' && v <= 'z' {
			for q := 0; q < 26; q++ {
				index := q
				if v == alphaLow[index] {
					runes[i] = alphaLow[index+14]
				}
			}
		} else if v >= 'A' && v <= 'Z' {
			for q := 0; q < 26; q++ {
				index := q
				if v == aplhaUp[index] {
					runes[i] = aplhaUp[index+14]
				}
			}
		}
	}

	changedStr := ""
	for _, r := range runes {
		changedStr += string(r)
	}

	return changedStr
}
