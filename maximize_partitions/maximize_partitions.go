package maximizepartitions

import (
	"fmt"
	"math/rand"
)

/*
You are given a string s and an integer k.

First, you are allowed to change at most one index in s to another lowercase English letter.

After that, do the following partitioning operation until s is empty:

Choose the longest prefix of s containing at most k distinct characters.
Delete the prefix from s and increase the number of partitions by one. The remaining characters (if any) in s maintain their initial order.
Return an integer denoting the maximum number of resulting partitions after the operations by optimally choosing at most one index to change.
*/
var charset = []rune("abcdefghijklmnopqrstuvwxyz")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func chooseLongestPrefix(s string, k int) int {
	distinctCharacters := map[rune]struct{}{}
	for i, c := range s {
		_, ok := distinctCharacters[c]
		if !ok && len(distinctCharacters) < k {
			distinctCharacters[c] = struct{}{}
		} else if !ok && len(distinctCharacters) >= k {
			return i
		}
	}
	return len(s)
}

func countPartitions(s string, k int) int {
	remaining := s
	partitions := 0
	for len(remaining) > 0 {
		i := chooseLongestPrefix(remaining, k)
		remaining = remaining[i:]
		partitions += 1
	}
	return partitions
}

func maxPartitionsAfterOperations(s string, k int) int {
	chars := map[rune]bool{}
	var newChar rune
	for _, c := range s {
		chars[c] = true
	}
	for i, c := range charset {
		if ok := chars[c]; !ok {
			newChar = c
			break
		}
		if i == 25 {
			newChar = rune(s[0]) + 3
		}
	}
	fmt.Printf("new char %s\n", string(newChar))
	numParts := 1
	for i := 1; i < len(s); i++ {
		sp := []rune(s)
		sp[i] = newChar
		p := countPartitions(string(sp), k)
		if p > numParts {
			numParts = p
		}
	}
	return numParts
}
