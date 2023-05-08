package main

import (
	"bytes"
	"fmt"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func addStrings(num1 string, num2 string) string {
	ret := ""
	result := bytes.NewBufferString(ret)
	i := 0
	j := 0
	var x byte
	for i < len(num1) || j < len(num2) {
		var first byte
		if i < len(num1) {
			first = num1[len(num1)-i-1] - '0'
			i++
		}
		var second byte
		if j < len(num2) {
			second = num2[len(num2)-j-1] - '0'
			j++
		}

		sum := first + second + x

		result.WriteByte('0' + sum%10)

		if sum >= 10 {
			x = 1
		} else {
			x = 0
		}
	}
	if x != 0 {
		result.WriteByte('0' + x)
	}
	return reverse(result.String())
}

func main() {
	fmt.Println(addStrings("123434563456", "123434563456"))
}
