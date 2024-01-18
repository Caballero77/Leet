package main

import (
	"fmt"
	"strings"
)

func formatLine(words []string, wordsLength int, maxWidth int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-wordsLength)
	}
	l := len(words)
	result := ""
	spaces := (maxWidth - wordsLength) / (l - 1)
	additionalSpace := (maxWidth - wordsLength) % (l - 1)
	fmt.Println(words, wordsLength, spaces, additionalSpace)
	for i := 0; i < l; i++ {
		result += words[i]
		if i < l-1 {
			result += strings.Repeat(" ", spaces)
			if i < additionalSpace {
				result += " "
			}
		}
	}
	return result
}

func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	lineStart := 0
	lineSize := 0
	for i := 0; i < len(words); i++ {
		if lineSize+(i-lineStart-1)+len(words[i]) >= maxWidth {
			result = append(result, formatLine(words[lineStart:i], lineSize, maxWidth))
			lineSize = 0
			lineStart = i
		}
		lineSize += len(words[i])
	}
	lastLine := strings.Join(words[lineStart:], " ")
	result = append(result, formatLine([]string{lastLine}, len(lastLine), maxWidth))
	return result
}

func main() {
	for _, line := range fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16) {
		fmt.Println(line, len(line))
	}
}
