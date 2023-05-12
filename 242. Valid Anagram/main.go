package main

import (
	"fmt"
	"unicode/utf8"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	first := make(map[rune]int)
	second := make(map[rune]int)
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		s = s[size:]
		if _, contains := first[r]; contains {
			first[r] += 1
		} else {
			first[r] = 1
		}
	}
	for len(t) > 0 {
		r, size := utf8.DecodeRuneInString(t)
		t = t[size:]
		if _, contains := second[r]; contains {
			second[r] += 1
		} else {
			second[r] = 1
		}
	}
	for k, firstV := range first {
		if secondV, contains := second[k]; !contains || secondV != firstV {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("abcd", "abcdeft"))
}
