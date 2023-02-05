package piscine

// func Abort(a, b, c, d, e int) int {
// 	array := [5]int{a, b, c, d, e}
// 	sum := 0
// 	for _, v := range array {
// 		sum += v
// 	}
// 	average := sum / len(array)
// 	return average
// }

func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array[2]
}
