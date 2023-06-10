package problem0058

import "strings"

func lengthOfLastWord(s string) int {
	trimmedStr := strings.TrimSpace(s)
	words := strings.Split(trimmedStr, " ")
	return len(words[len(words)-1])

}
