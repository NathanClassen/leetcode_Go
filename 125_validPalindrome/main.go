package main

import (
	"strings"
	"fmt"
)
/*

A phrase is a palindrome if, after converting all uppercase letters into lowercase
	letters and removing all non-alphanumeric characters, it reads the same forward and
	backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

 

Constraints:

    1 <= s.length <= 2 * 105
    s consists only of printable ASCII characters.


PEDAC


P

determine whether a string is a palindrome, considering only it's alphanumeric characters


E




D

A

use two pointers to start at either end of string and move toward middle
at each step, compare characters and if they're alphanumeric and equal, keep going
if they're AN and NOT equal, break
if one or both are NOT alphanumeric, ignore this comparison, and move the non-AN pointers one space in it's direction

set pointer `a` to 0
set pointer `b` to end of string
loop as long as a <= b
for each iteration
	if `a` and `b`
		point to alphn. and are equal
			a++
			b--
		point to alphn. and are !equal
			return false
		do not point at two alphan. chars
			if a is not alpha
				a++
			if b is not alpha
				b--
return true
		
*/

func main() {
	fmt.Println(isPalindrome("race car"))
	fmt.Println(isPalindrome("dad"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("r"))
	fmt.Println(isPalindrome("race, car"))
	fmt.Println(!isPalindrome("race cars"))
	fmt.Println(isPalindrome("ab_a"))
}

func isAlphaNumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
	a, b := 0, (len(s) - 1)

	for a <= b {
		ac := s[a]
		bc := s[b]
		acia := isAlphaNumeric(ac)
		bcia := isAlphaNumeric(bc)
		if acia && bcia {
			if ac == bc {
				a++
				b--
			} else {
				return false
			}
		} else {
			if !acia {
				a++
			}
			if !bcia {
				b--
			}
		}
	}
	return true
}