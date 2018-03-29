func reverseWords(s string) string {
	var buf string
	result := ""
    for i := 0; i < len(s); i++ {
		if (string(s[i]) == " ") {
			for j := len(buf) - 1; j >= 0; j-- {
				result += string(buf[j])
			}
			
			result += " "
			buf = ""
		} else {
			buf += string(s[i])
		}
	}
	for j := len(buf) - 1; j >= 0; j-- {
		result += string(buf[j])
	}
	return result
}