package main

import "fmt"

func isAnagram(s string, t string) bool {
	m := make(map[rune]int, len(s))

	for _, ch := range s {
		m[ch]++
	}

	for _, ch := range t {
		if m[ch] == 0 {
			return false
		}
		m[ch]--
		if m[ch] == 0 {
			delete(m, ch)
		}
	}

	if len(m) > 0 {
		return false
	} else {
		return true
	}

}

func main() {
	s := "abc"
	t := "cbad"
	res := isAnagram(s, t)
	fmt.Print(res)
}
