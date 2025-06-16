package main

import "fmt"

func firstUniqChar(s string) int {
	checked := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if checked[s[i]] {
			continue
		}
		uniq := true
		for j := i + 1; j < len(s); j++ {
			if checked[s[j]] {
				continue
			}
			if s[i] == s[j] {
				uniq = false
				checked[s[j]] = true
			}
		}
		if uniq {
			return i
		}
	}
	return -1
}

func main() {
	s := "cc"
	pos := firstUniqChar(s)

	fmt.Print(pos)
}
