package hangman

import "regexp"

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func IsAlpha(str string) bool {
	var alphanumeric = regexp.MustCompile("^[a-zA-Z]*$")
	return alphanumeric.MatchString(str)
}
