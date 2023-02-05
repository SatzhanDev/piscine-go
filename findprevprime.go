package piscine

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for i := nb; ; i-- {
		if IsPrime(i) == true {
			return i
		}
	}
}

// func IsPrime(nb int) bool {
// 	if nb < 2 {
// 		return false
// 	}
// 	for i := 2; i < nb; i++ {
// 		if nb%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
