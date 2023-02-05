package piscine

func Map(f func(int) bool, a []int) []bool {
	boolArr := make([]bool, len(a))
	for i := range a {
		boolArr[i] = f(a[i])
	}
	return boolArr
}
