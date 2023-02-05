package piscine

func SortWordArr(array []string) {
	for i := range array {
		for j := range array {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
