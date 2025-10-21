package longestcommonprefix

/*
14.
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

func longestCommonPrefix(strs []string) string {
	longest := 0
	for i := range len(strs[0]) {
		c := strs[0][i]
		match := true
		for j := range strs {
			if i > len(strs[j])-1 {
				match = false
				break
			}
			if strs[j][i] != c {
				match = false
				break
			}
		}
		if match {
			longest = i + 1
		} else {
			break
		}
	}
	return strs[0][0:longest]
}
