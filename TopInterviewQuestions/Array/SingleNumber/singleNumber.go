package main

import "fmt"

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	arr := []int{1, 1, 2, 4, 3, 2, 3}

	num := singleNumber(arr)
	fmt.Print(num)
}
