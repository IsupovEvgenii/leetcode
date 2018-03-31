/*Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".*/
func reverseString(s string) string {
	result := ""
	if s != "" {
		for pos := len(s) - 1; pos > len(s) / 2; pos-- {
			result += string(s[pos])
		}
		for pos := len(s) / 2; pos >= 0; pos-- {
			result += string(s[pos])
		}
	}
    
	return result
}
