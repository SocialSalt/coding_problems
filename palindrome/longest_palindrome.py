# def longestPalindrome(s: str):
#   gap = 0
#   longest = ''
#   for j in range(len(s)):
#     i = 0
#     while j-i >= gap:
#       substr = s[i:j+1]
#       if substr == substr[::-1]:
#         longest = substr
#         gap = len(longest)
#       i += 1
#   return longest


def longestPalindrome(s: str):
    gap = 0
    longest = ""
    n = len(s)
    for i in range(len(s)):
        r, l = i, i
        while l >= 0 and r < n and s[l] == s[r]:
            if r - l + 1 > gap:
                longest = s[l : r + 1]
                gap = r - l + 1
            l -= 1
            r += 1

    for i in range(len(s)):
        r, l = i + 1, i
        while l >= 0 and r < n and s[l] == s[r]:
            if r - l + 1 > gap:
                longest = s[l : r + 1]
                gap = r - l + 1
            l -= 1
            r += 1

    return longest


print(longestPalindrome("babad"))
print(longestPalindrome("a"))
print(longestPalindrome("abbc"))
