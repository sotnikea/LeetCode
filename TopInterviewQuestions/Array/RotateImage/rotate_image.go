package main

import "fmt"

func rotate(matrix [][]int) {
	for level := 0; level <= len(matrix)/2; level++ {
		for i := level; i < level+(len(matrix)-level*2)-1; i++ {
			up := &matrix[level][i]
			right := &matrix[i][len(matrix)-level-1]
			down := &matrix[len(matrix)-level-1][len(matrix)-i-1]
			left := &matrix[len(matrix)-i-1][level]

			*up, *right, *down, *left = *left, *up, *right, *down
		}
	}
}

func main() {
	matrix := [][]int{
		{5, 1},
		{2, 4},
	}

	rotate(matrix)
	fmt.Print(matrix)
}
