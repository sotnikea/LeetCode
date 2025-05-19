package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	allNums := make(map[int][]int)
	for i, val := range nums {
		allNums[val] = append(allNums[val], i)
	}
	var num1, num2 int
	for i, val := range nums {
		n := target - val
		if len(allNums[n]) > 0 {
			for j, index := range allNums[n] {
				if i == allNums[n][j] {
					continue
				}
				num1 = i
				num2 = index
			}
			if num1 != 0 {
				break
			}
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
