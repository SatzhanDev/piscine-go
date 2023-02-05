package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	newslice := [][]int{}
	sliceLen := len(slice)
	index := 0
	for i := index; i < sliceLen; i++ {
		if len(slice[index:]) >= size {
			newslice = append(newslice, slice[index:index+size])
			index += size
		} else {
			if len(slice[index:]) >= 1 {
				newslice = append(newslice, slice[index:])
				break
			} else {
				break
			}
		}
	}
	fmt.Println(newslice)
}
