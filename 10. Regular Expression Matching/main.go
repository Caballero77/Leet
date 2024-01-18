package main

import "fmt"

func isMatch(s string, p string) bool {
	if len(p) == 0 && len(s) != 0 {
		return false
	}
	if len(s) == 0 {
		if len(p) == 0 {
			return true
		}
		if p[0] == '*' {
			return isMatch(s, p[1:])
		}
		return false
	}
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if p[0] == '*' {
		return isMatch(s[1:], p[1:]) || isMatch(s[1:], p) || isMatch(s, p[1:])
	}
	if s[0] == p[0] || p[0] == '.' {
		return isMatch(s[1:], p[1:]) || isMatch(s, p[1:])
	}
	return isMatch(s, p[1:])
}

func main() {

	tests := [][]interface{}{
		{[]string{"aab", "c*a*b"}, true},
		{[]string{"mississippi", "mis*is*p*."}, false},
		{[]string{"aaa", "ab*ac*a"}, true},
		{[]string{"aaa", "ab*a*c*a"}, true},
		{[]string{"aaca", "ab*a*c*a"}, true},
		{[]string{"a", "ab*"}, true},
		{[]string{"bbbba", ".*a*a"}, true},
		{[]string{"ab", ".*.."}, true},
		{[]string{"ab", ".*..c*"}, true},
		{[]string{"a", ".*."}, true},
		{[]string{"aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s"}, true},
		{[]string{"abbbcd", "ab*bbbcd"}, true},
		{[]string{"bbab", "b*a*"}, false},
		{[]string{"a", "c*."}, true},
		{[]string{"a", "c*a"}, true},
		{[]string{"b", "a*."}, true},
		{[]string{"a", ".*a*"}, true},
		{[]string{"a", "..*"}, true},
		{[]string{"aabcbcbcaccbcaabc", ".*a*aa*.*b*.c*.*a*"}, true},
		{[]string{"abbabaaaaaaacaa", "a*.*b.a.*c*b*a*c*"}, true},
		{[]string{"bcaccbbacbcbcab", "b*.c*..*.b*b*.*c*"}, true},
		{[]string{"baabbbaccbccacacc", "c*..b*a*a.*a..*c"}, true},
		{[]string{"abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*b*"}, true},
		{[]string{"cbaacacaaccbaabcb", "c*b*b*.*ac*.*bc*a*"}, true},
		{[]string{"cbaacacaaccbaabcb", "c*b*b*.*ac*.*bc*a*"}, true},
		{[]string{"cabbbbcbcacbabc", ".*b.*.ab*.*b*a*c"}, true},
		{[]string{"abbcacbbbbbabcbaca", "a*a*.*a*.*a*.b*a*"}, true},
		{[]string{"aababbabacaabacbbbc", ".b*ac*.*c*a*b*.*"}, true},
		{[]string{"aaabaaaababcbccbaab", "c*c*.*c*a*..*c*"}, true},
		{[]string{"cbccaababcbabac", "c*aab*.*b.b.*.*a*."}, false},
		{[]string{"caccccaccbabbcb", "c*c*b*a*.*c*.a*a*a*"}, true},
		{[]string{"bbbaccbbbaababbac", ".b*b*.*...*.*c*."}, true},
		{[]string{"ccbbcabcbbaabaccc", "c*a*.*a*a*.*c*b*b*."}, true},
		{[]string{"abbaaaabaabbcba", "a*.*ba.*c*..a*.a*."}, true},
		{[]string{"bbcacbabbcbaaccabc", "b*a*a*.c*bb*b*.*.*"}, true},
		{[]string{"aabccbcbacabaab", ".*c*a*b.*a*ba*bb*"}, true},
		{[]string{"cbbbaccbcacbcca", "b*.*b*a*.a*b*.a*"}, true},
		{[]string{"cbacbbabbcaabbb", "b*c*.*a*..a.*c*.*"}, true},
		{[]string{"abaabababbcbcabbcbc", "b*ab.*.*.*.b..*"}, true},
		{[]string{"caaacccbaababbb", "c*.*b*ba*ac*c*b*.*"}, true},
		{[]string{"abbbaabccbaabacab", "ab*b*b*bc*ac*.*bb*"}, true},
		{[]string{"abbbaabccbaabacab", "ab*b*b*bc*ac*.*bb*"}, true},
		{[]string{"cacbacccbbbccab", ".b*b*.*c*a*.*bb*"}, true},
		{[]string{"abcbccbcbaabbcbb", "c*a.*ab*.*ab*a*..b*"}, true},
		{[]string{"caabbabbbbccccbbbcc", ".b*c*.*.*bb*.*.*"}, true},
		{[]string{"caaccabbbabcacaac", "b*c*b*b*.b*.*c*a*c"}, true},
		{[]string{"cbcaabcbaabccbaa", "c*b*ab*.*b*c*a*"}, false},
		{[]string{"bccbcccbcbbbcbb", "c*c*c*c*c*.*.*b*b*"}, true},
		{[]string{"ccacbcbcccabbab", ".c*a*aa*b*.*b*.*"}, true},
		{[]string{"aabbcbcacbacaaccacc", "c*b*b*.*.*.*a*.*"}, true},
		{[]string{"bcbabcaacacbcabac", "a*c*a*b*.*aa*c*a*a*"}, true},
		{[]string{"acabbabacaccacccabc", "a*.*c*a*.b.*a*.*"}, true},
		{[]string{"babbcccbacaabcbac", "b.*.*c*b*b*.*c*c"}, true},
		{[]string{"cbbbbabaabbacbbc", "a*c*b*.*bb*a*.*a*"}, true},
		{[]string{"accbabbacbbbacb", ".*.*.*a*bba*ba*"}, false},
		{[]string{"ababbcaaabbaccb", "c*c*..*a*a*a*.*"}, true},
		{[]string{"bcabcbcaccabcbb", "a*a*c*a*.*a*c*bc*."}, true},
		{[]string{"bcbbbacbabccbabbac", "c*.*b*a.*a*a*a*"}, true},
		{[]string{"ccbbbbbacacaaabcaa", ".*ba*.*.b*c*c*b*a.*"}, true},
		{[]string{"acaababbccbaacabcab", "..*bb*b*c*a*c*.*.b"}, true},
		{[]string{"cbabcabbbacbcaca", "a*c*.*a*a*b*c*a*.*"}, true},
		{[]string{"bacacaababbbcbc", ".*a*.*a*.aa*c*b*c"}, false},
		{[]string{"cbabcbbaabbcaca", ".a*b*.*.*b*c*.*b*a*"}, true},
		{[]string{"bbaaaacabccbcac", "b*b*a*c*c*a*c*.*"}, true},
		{[]string{"bcccccbaccccacaa", ".*bb*c*a*b*.*b*b*c*"}, true},
		{[]string{"bcbaccbbbccabaac", "c*.*a*b*ac*a*a*"}, true},
		{[]string{"bacacbacaaabccbcbaa", "a*.c*c*c*a*b*..*"}, true},
		{[]string{"baccbbcbcacacbbc", "c*.*b*c*ba*b*b*.a*"}, true},
	}
	for _, c := range tests {
		arr := c[0].([]string)
		fmt.Println(c[0], c[1] == isMatch(arr[0], arr[1]))
	}

	fmt.Println(isMatch("a", ".*a*"))
}
