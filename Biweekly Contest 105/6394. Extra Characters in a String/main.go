package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minExtraChar(s string, dictionary []string) int {
	dict := make(map[string]bool)
	for _, word := range dictionary {
		dict[word] = true
	}
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if i > 0 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		for j := 0; j <= i; j++ {
			t := s[j : i+1]
			if ok, _ := dict[t]; ok {
				if j == 0 {
					dp[i] = 0
				} else {
					dp[i] = min(dp[i], dp[j-1])
				}
			}
		}
	}
	return dp[len(s)-1]
}

func main() {
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code"}))
	fmt.Println(minExtraChar("kevlplxozaizdhxoimmraiakbak", []string{"yv", "bmab", "hv", "bnsll", "mra", "jjqf", "g", "aiyzi", "ip", "pfctr", "flr", "ybbcl", "biu", "ke", "lpl", "iak", "pirua", "ilhqd", "zdhx", "fux", "xaw", "pdfvt", "xf", "t", "wq", "r", "cgmud", "aokas", "xv", "jf", "cyys", "wcaz", "rvegf", "ysg", "xo", "uwb", "lw", "okgk", "vbmi", "v", "mvo", "fxyx", "ad", "e"}))
}
