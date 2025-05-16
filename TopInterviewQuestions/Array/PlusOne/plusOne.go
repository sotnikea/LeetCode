package main

import "fmt"

func plusOne(digits []int) []int {
	var digitsCopy []int
	digitsCopy = append(digitsCopy, digits...)
	for i := len(digitsCopy) - 1; i >= 0; i-- {
		if digitsCopy[i] == 9 {
			digitsCopy[i] = 0
			if i == 0 {
				var res []int
				res = append(res, 1)
				res = append(res, digitsCopy...)
				return res
			}
		} else {
			digitsCopy[i]++
			break
		}
	}
	return digitsCopy
}

func main() {
	nums := []int{9, 9}
	resNums := plusOne(nums)
	fmt.Print(resNums)
}
