package main

import (
	"fmt"
)

func isApthanumeral(l byte) bool {
	return (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') || (l >= '0' && l <= '9')
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
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func partition(s string) [][]string {
	result := [][]string{}
	var inner func(string, int, []string)
	inner = func(s string, index int, curr []string) {
		if index == len(s) {
			cp := make([]string, len(curr))
			copy(cp, curr)
			result = append(result, cp)
			return
		}
		for i := index + 1; i <= len(s); i++ {
			if isPalindrome(s[index:i]) {
				curr = append(curr, s[index:i])
				inner(s, i, curr)
				curr = curr[:len(curr)-1]
			}
		}
	}

	inner(s, 0, []string{})
	return result
}

func main() {
	fmt.Println(partition("bdhgrsdssskft"))
}
