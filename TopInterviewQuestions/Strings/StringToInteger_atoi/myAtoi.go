package main

import "fmt"

func myAtoi(s string) int {
	var num int64
	num = 0
	isNegative := false
	isPrechecked := false

	for i, ch := range s {
		// Whitespace
		if ch == ' ' {
			if !isPrechecked {
				continue
			} else {
				break
			}
		}

		// Signedness
		if !isPrechecked {
			isPrechecked = true
			if ch == '-' {
				isNegative = true
				continue
			} else if ch == '+' {
				continue
			}
		}

		// Conversion
		if ch >= 48 && ch <= 57 {
			num = num*10 + int64(int(ch)-48)
			if num > 2147483647 {
				break
			}
		} else {
			break
		}

		if i < len(s)-1 {
			continue
		}
	}

	// Rounding
	if isNegative {
		num = -num
	}
	if num < -2147483648 {
		num = -2147483648
	}
	if num > 2147483647 {
		num = 2147483647
	}

	return int(num)
}

func main() {
	s := "9223372036854775808"
	num := myAtoi(s)
	fmt.Print(num)
}
