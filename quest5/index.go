package piscine

// func Index(s string, toFind string) int {
// 	str1 := []rune(s)
// 	str2 := []rune(toFind)
// 	count1 := 0
// 	count2 := 0
// 	for index := range str2 {
// 		index = index
// 		count2++
// 	}
// 	if count2 == 0 {
// 		return 0
// 	}
// 	for index := range str1 {
// 		index = index
// 		count1++
// 	}
// 	for index, letter := range str1 {
// 		if letter == str2[0] && count1 >= count2+index-1 {
// 			a := 1
// 			for b := 1; b < count2; b++ {
// 				if str2[b] == str1[index+b] {
// 					a++
// 				}
// 			}
// 			if a == count2 {
// 				return index
// 			}
// 		}
// 	}
// 	return -1
// }

func Index(s string, toFind string) int {
	length := 0
	sublength := 0

	for i := range s {
		length = i + 1
	}

	for i := range toFind {
		sublength = i + 1
	}

	t := 0

	for i := 0; i < length; i++ {
		j := 0
		t = i
		for j < sublength {
			if t < length && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == sublength {
			return i
		}

	}
	return -1
}
