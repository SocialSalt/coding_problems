package dynamicprogramming

import "math"

/*
1653.  Minimum Deletions to Make String Balanced
You are given a string s consisting only of characters 'a' and 'b'.

You can delete any number of characters in s to make s balanced. s is balanced if there is no pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.

Return the minimum number of deletions needed to make s balanced.
*/

func minimumDeletions(s string) int {
	// approach:
	// scores[i] is the number of b's after i and the number of a's before i if we
	// delete all the b's before i and the a's after i then the string is balanced.
	// So if we get the min of scores then that's the minimum number of deletions
	// need to make the string balanced.

	aCount := 0
	bCount := 0
	scores := make([]int, len(s))

	for i := range s {
		scores[i] = scores[i] + bCount
		if s[i] == 'b' {
			bCount++
		}
		j := len(s) - 1 - i
		scores[j] = scores[j] + aCount
		if s[j] == 'a' {
			aCount++
		}
	}

	minDel := math.MaxInt
	for i := range scores {
		if scores[i] < minDel {
			minDel = scores[i]
		}
	}
	return minDel
}

func minimumDeletionsSimple(s string) int {
	// approach:
	// Consider the pairs of "a" after a "b" or "b" before an "a". It doesn't
	// actually matter if the delete the "a" or the "b", we just know that one of
	// them needs to be deleted. Because of this we can just count the number of
	// times we encounter such a pair. Every time we see an "a" after seeing a "b",
	// there must be some deletion. We count that deletion and then remove that "b"
	// count, because that pair has been resolved.

	dels := 0
	bCount := 0
	for i := range s {
		if s[i] == 'a' && bCount > 0 {
			bCount--
			dels++
		}
		if s[i] == 'b' {
			bCount++
		}
	}

	return dels
}
