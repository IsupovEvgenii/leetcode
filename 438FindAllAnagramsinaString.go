/*Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".*/
package main

import (
	"fmt"
)

func getPattern(p string) map[string]int {
	patternMap := make(map[string]int)
	for _, v := range p {
		patternMap[string(v)] += 1
	}
	return patternMap
}

func checkStr(str string, pat string) bool {
	pattern := getPattern(pat)
	number1 := 0
	for i := 0; i < len(str); i++ {
		if _, ok := range pattern[string(str[i])]; ok {
			pattern[string(str[i])]--
			if pattern[v] < 0 {
				return 0
			}
			if pattern[v] == 0 {
				number1++
			}
			
		}
	}
	if number1 == len(pattern) {
		return 1
	} else {
		return 3
	}
}

// func checkPattern(pattern map[string]int) int {
// 	number1 := 0
	
// 	for v := range pattern {
// 		if pattern[v] < 0 {
// 			return 0
// 		}
// 		if pattern[v] == 0 {
// 			number1++
// 		}
// 	}
// 	if number1 == len(pattern) {
// 		return 1
// 	} else {
// 		return 3
// 	}
// }

// func findAnagrams(s string, p string) []int {
// 	var result []int

// 	for i := 0; i < len(s); i++ {
// 		dict := getPattern(p)

// 		if _, ok := dict[string(s[i])]; ok {

// 			dict[string(s[i])] -= 1

// 			if checkPattern(dict) == 1 {

// 				result = append(result, i)

// 			} else {

// 				for j := i + 1; j < len(s); j++ {

// 					if _, ok := dict[string(s[j])]; ok {

// 						dict[string(s[j])] -= 1

// 						if checkPattern(dict) == 1 {

// 							result = append(result, i)
// 							break

// 						}
// 						if checkPattern(dict) == 0 {

// 							break

// 						}

// 					} else {
// 						break
// 					}
// 				}

// 			}

// 		}
// 	}

// 	return result

// }

func findAnagrams(s string, p string) []int {
	var result []int
	dict := getPattern(p)
	iter1 := 0
	str1 := ""
	array1 := []int{}
	for i:= 0; i < len(s); i++ {
		if _, v := range dict[string(s[i])] {
			str1 = string(s[i]) + str1
			array1 = append(array1, i)
			if checkStr(str1, p) == 1 {
				return 
			}
		}
	}

	return result
 
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
