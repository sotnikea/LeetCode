package main

import (
	"fmt"
)

func rotate1(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	tmpArr := make([]int, k)

	// Save k elements
	for i := 0; i < k; i++ {
		tmpArr[i] = nums[len(nums)-1-i]
	}

	// Change element position
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}

	// Set first k element from tmp array
	for i := 0; i < k; i++ {
		nums[i] = tmpArr[k-i-1]
	}
}

func rotate2(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	for i := 0; i < k; i++ {

		j := len(nums) - 1
		tmp := nums[j]

		for ; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = tmp
	}
}

func rotate3(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	tmpArr := make([]int, len(nums))

	tmpPos := 0
	for i := len(nums) - k; i < len(nums); i++ {
		tmpArr[tmpPos] = nums[i]
		tmpPos++
	}

	for i := 0; i < len(nums)-k; i++ {
		tmpArr[tmpPos] = nums[i]
		tmpPos++
	}

	copy(nums, tmpArr)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	num := 3

	rotate3(arr, num)

	fmt.Print(arr)
}
