package piscine

func FindNextPrime(nb int) int {
	for i := nb; ; i++ {
		if IsPrime(i) == true {
			return i
		}
	}
}
