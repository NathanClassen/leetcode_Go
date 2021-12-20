/*
	Given a string `s`, find the length of the longest substring without repeating characters.


	Example 1:

	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

	Example 2:

	Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

	Example 3:

	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

	Constraints:

    0 <= s.length <= 5 * 104
    s consists of English letters, digits, symbols and spaces.
	___________________________________________________________________________________________

	- looking for longest sequence of chars, where current char has not been seen before
	- gotchas:
		- we may find a sequence of 4 chars starting from the first char, but then may find a unique sequence of 5 starting from the 2nd or 3rd char, etc
			- so, at first glance, seems we will need to examine entire string (breaking from examination upon first repeated char) starting from each char
			- can stop looking if `currentLongest` >= remaining characters in string from current starting place

	- will need a way to track `seen` chars
		- map of map[string]bool seems a good choice

		would have to somehow "clear" the map after each "outer" iteration...
		alternatively, could keep track of an int value with each outer iteration `0` on one iteration, then `1` on the next, and so on, so that "tuthiness" is held as "whatever the current"
		`int` tracker is. This would clearing un-necessary. Or could just rest map to empty map... certainly simpler and less space
	- will have to track longest, and a current longest variable
		- at the end of a partial iteration, only replace longest if current longest is greater

	Algo:

		general outline:
			- loop while tracking to pointers
			- 1st theme is recording, eval, processing
			- 2nd theme is loop status control via pointer manipulation
				if `criteria` r = n
				if `criteria` l = x
				if `criteria` break/return


		set all variables
			- l int 0
			- r int 0
			- localLongest int
			- globalLongest int
			- seen map[byte]true

		for
			see char at r
				if char not seen
					add to seen
					incr localL
					incr r
				if char seen
					clear seen
					maybe replace globalL
					clear localL
					l++
					r = l

				if length of globalL is larger than remaining chars
					break

				if l >= len(s)
					break

*/

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("aacabczc"))
}

func lengthOfLongestSubstring(s string) int {
	l := 0
	r := 0
	localL := 0
	globalL := 0
	seen := map[byte]bool{}
	sLen := len(s)

	if sLen == 0 {
		return 0
	}

	for {
		char := s[r]
		if _, haveSeen := seen[char]; haveSeen {
			seen = map[byte]bool{}
			if localL > globalL {
				globalL = localL
			}
			localL = 0
			l++
			r = l

			if globalL >= sLen-l {
				break
			}
		} else {
			seen[char] = true
			localL++
			r++
		}

		if l >= sLen || r >= sLen {
			break
		}
	}

	if globalL > localL {
		return globalL
	} else {
		return localL
	}
}
