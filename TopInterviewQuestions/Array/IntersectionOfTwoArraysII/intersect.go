package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	res := make(map[int]int)
	for _, num1 := range nums1 {
		res[num1]++
	}

	var resArr []int
	for _, num2 := range nums2 {
		if val, ok := res[num2]; ok && val > 0 {
			res[num2]--
			resArr = append(resArr, num2)
		}
	}

	return resArr
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	res := intersect(nums1, nums2)
	fmt.Print(res)
}
