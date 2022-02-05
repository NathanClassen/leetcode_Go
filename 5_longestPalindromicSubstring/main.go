/*
	Given a string s, return the longest palindromic substring in s.

	Example 1:
	Input: s = "badac"
	Output: "bab"
	Explanation: "aba" is also a valid answer.

	Example 2:
	Input: s = "cbbd"
	Output: "bb"

	Constraints:
		1 <= s.length <= 1000
		s consist of only digits and English letters.

	- check for all palindromes, tracking the longest one
	- check for palindromes
		- does a given substr == substring.reversed?
	- check for all
		- the longest palindrome, or any palindrome, could start at any given index of the string
		- need to check each location of the string (starting index) for palindromes, tracking the longest found


	isPalindrome?
		set idx at start and end
			move toward each other, ensuring that the characters at each index match
			stop loop when after comparison, both indeces are equal or consecutive
		return true after loop
		if ever the chars do not match in the loop, return false

	find palindromes
		for each index, check index + endindex string to see if palindrome, if so, save it as longest if it's longest
		set idx at 0 and ridx at idx, send string to check
		increment ridx and check that string,
			repeat until ridx has reached end of string
			then increment idx and repeat
				do this until idx has reached end of string, or until remaining length of unchecked string is shorter than longest palindrom found
*/

package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("roitya"))
}

func longestPalindrome(s string) string {
	longest := ""
	if len(s) == 1 {
		return s
	}
	for idx := 0; idx < len(s); idx++ {
		for ridx := idx + 1; ridx <= len(s); ridx++ {
			substr := s[idx:ridx]
			if isPalindrome(substr) {
				if len(substr) > len(longest) {
					longest = substr
				}
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	lIdx, rIdx := 0, len(s)-1
	for {
		if s[lIdx] != s[rIdx] {
			return false
		}
		if lIdx == rIdx || rIdx-lIdx == 1 {
			return true
		}
		lIdx++
		rIdx--
	}
}
