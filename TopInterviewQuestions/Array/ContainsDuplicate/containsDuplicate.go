package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	valMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if val, ok := valMap[nums[i]]; ok {
			valMap[nums[i]] = val + 1
		} else {
			valMap[nums[i]] = 1
		}
	}

	for _, val := range valMap {
		if val > 1 {
			return true
		}
	}

	return false
}

func main() {
	arr := []int{1, 2, 3, 1}

	res := containsDuplicate(arr)
	fmt.Println(res)
}
