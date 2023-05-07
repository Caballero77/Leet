package main

import "fmt"

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	return result
}

func isPalindrome(x int) bool {
	return x >= 0 && x == reverse(x)
}

func main() {
	fmt.Println(reverse(-102))
	fmt.Println(isPalindrome(121))
}
