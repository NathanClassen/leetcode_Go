/*
	Given a string s, find the length of the longest substring without repeating characters.

	Example 1:
                l
	Input: s = "abcabcbb"
                r

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
*/
/*
	Problem:
	find the longest string of unique characters within a string
	basically need to check from each position of the string
	if there is only one char in str, thats the longest
	if no chars, 0 is the answer

	Mental model:
	track the longest sequence found, as you iterate string from each character as a starting position,
	checking how many consecutive chars from each starting position are unique

	Datastructure:
	iteration and length tracking can be done with a growing and sliding window
	may use map to track seen chars

	algorithm:
	set l, r to 0
	set longest to 0
	as long as `l` is less than len(str)
		if r is >= len(str)
			l++
			r = l
			globalL = max(longest, globalL)
		check if r char has been seen
			if not
				++longest
				++r
			if so
				l++
				r = l
			globalL = max(longest, globalL)

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("asdfghjklqwertyuiopzxcvbnm"))
}

func lengthOfLongestSubstring(s string) int {
	globL := 0
	currL := 0
	seen := map[byte]bool{}
	l, r := 0, 0

	for l < len(s) && len(s)-l > globL {
		if r >= len(s) || seen[s[r]] {
			l++
			r = l
			globL = int(math.Max(float64(currL), float64(globL)))
			currL = 0
			seen = map[byte]bool{}
			continue
		}
		seen[s[r]] = true
		currL++
		r++
	}
	return globL
}

// abca
