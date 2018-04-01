/*Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.*/
func in_row1(a string) bool {
	row1 := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"}
	for _, b := range row1 {
		if b == a {
			return true
		}
	}
	return false
}
func in_row2(a string) bool {
	row2 := []string{"a", "s", "d", "f", "g", "h", "j", "k", "l", "A", "S", "D", "F", "G", "H", "J", "K", "L"}
	for _, b := range row2 {
		if b == a {
			return true
		}
	}
	return false
}
func in_row3(a string) bool {
	row3 := []string{"z", "x", "c", "v", "b", "n", "m", "Z", "X", "C", "V", "B", "N", "M"}
	for _, b := range row3 {
		if b == a {
			return true
		}
	}
	return false
}

func findWords(words []string) []string {
	var answer []string
	for i := 0; i < len(words); i++ {
		result := true
		checker := 0
		for j := 0; j < len(words[i]); j++ {
			if checker == 0 {
				if in_row1(string(words[i][j])) {
					checker = 1
				}
				if in_row2(string(words[i][j])) {
					checker = 2
				}
				if in_row3(string(words[i][j])) {
					checker = 3
				}
			} else {
				if in_row1(string(words[i][j])) {
					if checker != 1 {
						result = false
						break
					}

				}
				if in_row2(string(words[i][j])) {
					if checker != 2 {
						result = false
						break
					}
				}
				if in_row3(string(words[i][j])) {
					if checker != 3 {
						result = false
						break
					}
				}
			}

		}
		if result == true {
			answer = append(answer, words[i])
		}
	}
	return answer
}

