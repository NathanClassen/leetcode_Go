package main

/*
	Given two strings s and t, return true if t is an anagram of s, and false otherwise.

	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

	

	Example 1:

	Input: s = "anagram", t = "nagaram"
	Output: true

	Example 2:

	Input: s = "rat", t = "car"
	Output: false

	

	Constraints:

		1 <= s.length, t.length <= 5 * 104
		s and t consist of lowercase English letters.

	

	Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

	approach 1:

	sort both strings and compare the result
	O(logn) time
	
	approach 2:
	two pointers
	O(n2)

	approach 3:
	map
	iterate s adding each char to a map[rune]int incrementing the int if existing
	iterate t decrementing the int each time it's key comes up
	evaluate if all keys map to 0
	O(n) - time
	O(n) - space


*/

import "fmt"

func main() {
	fmt.Println(isAnagram("rat", "tar") == true)
	fmt.Println(isAnagram("cattaco", "tacocat") == true)
	fmt.Println(isAnagram("boom", "bomb") == false)
	fmt.Println(isAnagram("sniper", "pinser") == true)
	fmt.Println(isAnagram("", "") == true)
	fmt.Println(isAnagram("a", "a") == true)
	fmt.Println(isAnagram("a", "aa") == false)
}

func isAnagram(s string, t string) bool {
	var slen int
	if slen = len(s); slen != len(t) {
		return false
	}


	runes := make(map[rune]int, slen)

	for _, y := range s {
		if _, p := runes[y]; p {
			runes[y]++
		} else {
			runes[y] = 1
		}
	}

	for _, y := range t {
		if _, p := runes[y]; p {
			runes[y]--
		} else {
			return false
		}
	}

	for _, v := range runes {
		if v != 0 {
			return false
		}
	}

	return true
}