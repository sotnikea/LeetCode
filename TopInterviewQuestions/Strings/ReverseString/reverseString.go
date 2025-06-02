package main

import "fmt"

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < len(s)/2 {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	data := []byte{'1', '2', '3', '4', '5'}
	reverseString(data)
	fmt.Print(string(data))
}
