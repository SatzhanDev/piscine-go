package piscine

func TrimAtoi(s string) int {
	flag := false
	count := 0
	for _, a := range s {
		if a == '-' && count == 0 {
			flag = true
		}
		if a >= 48 && a <= 57 {
			count = count*10 + int(a-48)
		}
	}
	if flag {
		return -count
	} else {
		return count
	}
}
