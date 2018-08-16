/*
We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words.

You may return the list in any order.



Example 1:

Input: A = "this apple is sweet", B = "this apple is sour"
Output: ["sweet","sour"]
Example 2:

Input: A = "apple apple", B = "banana"
Output: ["banana"]*/
func uncommonFromSentences(A string, B string) []string {
	words := strings.Fields(A)
	words = append(words, strings.Fields(B)...)
	sort.Strings(words)
	var uncommonWords []string
	for i := 0; i < len(words)-1; i++ {
		if words[i] != words[i+1] {
			uncommonWords = append(uncommonWords, words[i])
		} else {
			for words[i] == words[i+1] && i+1 != len(words)-1 {
				i++
			}
		}
	}
	if words[len(words)-1] != words[len(words)-2] {
		uncommonWords = append(uncommonWords, words[len(words)-1])
	}
	return words
}
