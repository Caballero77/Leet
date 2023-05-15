package main

import (
	"fmt"
)

func isApthanumeral(l byte) bool {
	return (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') || (l >= '0' && l <= '9')
}

func toLowwer(l byte) byte {
	if l >= 'A' && l <= 'Z' {
		return (l - 'A') + 'a'
	}
	return l
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !isApthanumeral(s[i]) {
			i++
		}
		for i < j && !isApthanumeral(s[j]) {
			j--
		}
		if toLowwer(s[i]) != toLowwer(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("test   tset"))
}
