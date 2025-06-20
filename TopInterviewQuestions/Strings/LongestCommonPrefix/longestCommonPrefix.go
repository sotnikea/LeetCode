package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	pos := -1
	for i:=0;;i++ {
		letter := strs[0][i]
		for _, text
	}
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Print(res)
}
