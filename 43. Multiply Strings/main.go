package main

import (
	"bytes"
	"fmt"
)

func multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	result := make([]byte, len1+len2)
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			next := i + j
			current := i + j + 1

			mult := (num1[i]-'0')*(num2[j]-'0') + result[current]

			result[current] = mult % 10
			result[next] += mult / 10
		}
	}
	buffer := bytes.NewBufferString("")
	for i := 0; i < len(result); i++ {
		if result[i] != 0 || len(buffer.Bytes()) != 0 {
			buffer.WriteByte('0' + result[i])
		}
	}
	if len(buffer.Bytes()) == 0 {
		return "0"
	}
	return buffer.String()
}

func main() {
	fmt.Println(multiply("123", "456"))
}
