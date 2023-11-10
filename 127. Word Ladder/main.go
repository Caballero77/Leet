package main

import (
	"fmt"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func isNext(w1, w2 string) bool {
	x := true
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			if x {
				x = false
			} else {
				return false
			}
		}
	}
	return !x
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !contains(wordList, endWord) {
		return 0
	}
	set := make(map[string]bool)
	queue := make([]string, 0)

	queue = append(queue, beginWord, "")

	depth := 1
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == "" && len(queue) > 0 {
			depth++
			queue = append(queue, "")
			continue
		}
		if curr == endWord {
			return depth
		}
		for _, word := range wordList {
			if !set[word] && isNext(curr, word) {
				queue = append(queue, word)
				set[word] = true
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	//fmt.Println(ladderLength("hot", "dog", []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"}))
}
