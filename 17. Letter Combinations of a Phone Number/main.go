package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	numbers := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := make([]string, 0)

	var inner func(int, string)

	inner = func(index int, curr string) {
		if index == len(digits) {
			if curr != "" {
				result = append(result, curr)
			}
			return
		}

		letters := numbers[digits[index]]
		for i := 0; i < len(letters); i++ {
			inner(index+1, curr+letters[i:i+1])
		}
	}

	inner(0, "")
	return result
}

func main() {
	fmt.Println(letterCombinations("2"))
}
