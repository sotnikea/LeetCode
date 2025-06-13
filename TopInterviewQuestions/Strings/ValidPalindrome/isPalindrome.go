package main

import (
	"fmt"
	"slices"
	"unicode"
)

func isPalindrome(s string) bool {
	var forward []string
	var backward []string

	for _, ch := range s {
		if int(ch) >= 65 && int(ch) <= 90 {
			forward = append(forward, string(unicode.ToLower(ch)))
		}

		if int(ch) >= 48 && int(ch) <= 57 ||
			int(ch) >= 97 && int(ch) <= 122 {
			forward = append(forward, string(ch))
		}
	}

	for i := len(forward) - 1; i >= 0; i-- {
		backward = append(backward, forward[i])
	}

	if slices.Equal(forward, backward) {
		return true
	} else {
		return false
	}

}

func main() {
	s := "ab a"
	res := isPalindrome(s)
	fmt.Print(res)
}
