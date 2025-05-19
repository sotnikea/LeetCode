package main

import "fmt"

func twoSum(nums []int, target int) []int {
	allNums := make(map[int]int)
	for i, val := range nums {
		allNums[val] = i
	}
	var num1, num2 int
	for i, val := range nums {
		n = target - val
		if allNums[n] > 0 {
			num1 = i
			num2 = val
			break
		}
	}
	return []int{num1, num2}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Print(res)
}
