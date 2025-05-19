package main

import "fmt"

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
		if nums[i] == 0 {
			break
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Print(nums)
}
