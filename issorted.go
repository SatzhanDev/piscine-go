package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	cap := len(a)
	sort := 0
	for i := 0; i < len(a)-1; i++ {
		isSorted := f(a[i], a[i+1])
		if isSorted > 0 {
			sort++
		} else if isSorted < 0 {
			sort--
		} else {
			cap--
		}
	}
	cap--

	if sort == cap || sort == -1*cap {
		return true
	}
	return false
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
