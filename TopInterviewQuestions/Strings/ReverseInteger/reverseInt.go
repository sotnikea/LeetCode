package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	maxInt32 := math.Pow(2, 31) - 1
	minInt32 := math.Pow(-2, 31)

	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}

	if res < int(maxInt32) && res > int(minInt32) {
		return res
	} else {
		return 0
	}
}

func main() {
	num := -123
	res := reverse(num)
	fmt.Print(res)
}
