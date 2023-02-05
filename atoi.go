package piscine

func Atoi(nbr string) int64 {
	var res int64 = 0
	var sign int64 = 1
	if IsNumeric(nbr) {
		if nbr[0] == '-' {
			nbr = nbr[1:]
			sign *= -1
		} else {
			nbr = nbr[1:]
		}
		for _, digit := range nbr {
			tmp := int64(digit-48) * sign
			if AtoiOverflow(res, int64(10), tmp) {
				res = res*10 + tmp
			}
		}

	} else {
		return 0
	}
	return res
}

func AtoiOverflow(res, b, tmp int64) bool {
	if res < 0 && tmp < 0 {
		return res*b+tmp < 0
	} else if res > 0 && tmp > 0 {
		return res*b+tmp > 0
	}
	return true
}

// func IsNumeric(s string) bool {
// 	a := []rune(s)
// 	for _, v := range a {
// 		if (v < 48 || v > 57) && v != '-' {
// 			return false
// 		}
// 	}
// 	return true
// }
