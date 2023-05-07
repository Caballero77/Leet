package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		match := 0
		j := 0
		for j < len(needle) && i+j < len(haystack) && haystack[i+j] == needle[j] {
			match++
			j++
		}

		if match == len(needle) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("sasadbutsad", "sad"))
}
