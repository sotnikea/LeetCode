package main

import "fmt"

func removeDuplicates(nums []int) int {
	uniqIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[uniqIndex] != nums[i] {
			uniqIndex++
			nums[uniqIndex] = nums[i]
		}
	}
	return uniqIndex + 1
}

func main() {
	testNum := []int{1, 1, 1, 2, 3, 3, 3, 4, 4, 5, 6, 6, 7, 7, 7, 7, 7, 9, 10, 10, 11}
	expectedNums := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11}

	k := removeDuplicates(testNum)

	if k != len(expectedNums) {
		fmt.Print("Checking array size failed!")
		return
	}

	for i := 0; i < k; i++ {
		if testNum[i] != expectedNums[i] {
			fmt.Print("Checking array element failed!")
			return
		}
	}

	fmt.Print("Success!")
}
