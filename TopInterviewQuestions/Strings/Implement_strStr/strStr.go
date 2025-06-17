package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	pos := -1
	j := 0
	for {
		i := 0
		// Skip unmatched
		for ; ; j++ {
			if j >= len(haystack) {
				return -1
			}
			if haystack[j] == needle[i] {
				break
			}
		}

		pos = j
		for {
			if i >= len(needle) {
				return pos
			}
			if j >= len(haystack) {
				return -1
			}
			if haystack[j] == needle[i] {
				i++
				j++
			} else {
				j = pos + 1
				break
			}
		}
	}
}

func main() {
	haystack := "leetcode"
	needle := "leeto"

	res := strStr(haystack, needle)
	fmt.Print(res)
}
