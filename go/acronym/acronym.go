// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate returns the first letter in uppercase of the given words.
func Abbreviate(s string) string {
	var acronym strings.Builder
	startOfWord := false
	for index, letter := range s {
		if index == 0 {
			acronym.WriteRune(letter)
		} else if startOfWord && ((letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z')) {
			startOfWord = false
			if letter >= 'a' && letter <= 'z' {
				acronym.WriteRune(letter - 'a' + 'A')
			} else {
				acronym.WriteRune(letter)
			}
		} else if !startOfWord && !((letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z')) && letter != '\'' {
			startOfWord = true
		}
	}
	return acronym.String()
}
