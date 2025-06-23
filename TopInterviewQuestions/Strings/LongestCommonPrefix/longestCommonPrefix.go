package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	pos := -1
	for j, ch := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= j || ch != rune(strs[i][j]) {
				if pos == -1 {
					return ""
				} else {
					return strs[0][0 : pos+1]
				}
			}
		}
		pos = j
	}
	return strs[0][0 : pos+1]
}

func main() {
	strs := []string{"fl", "fl", "fl"}
	res := longestCommonPrefix(strs)
	fmt.Print(res)
}
