package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	var targ []int
	for ind1, val1 := range nums {
		for ind2, val2 := range nums {
			if val1+val2 == target {
				targ = append(targ, ind1, ind2)
				return targ
			}
		}
		break
	}
	return nil
}

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}
