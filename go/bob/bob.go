// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

// Hey responses as per the given remark
func Hey(remark string) string {
	switch {
	case IsSilence(remark):
		return "Fine. Be that way!"
	case IsAQuestion(remark) && !IsAllUppercase(remark):
		return "Sure."
	case !IsAQuestion(remark) && IsAllUppercase(remark):
		return "Whoa, chill out!"
	case IsAQuestion(remark) && IsAllUppercase(remark):
		return "Calm down, I know what I'm doing!"
	default:
		return "Whatever."
	}
}

// IsAQuestion returns true if the string ends with '?'
func IsAQuestion(remark string) bool {
loop:
	for i := len(remark) - 1; i > 0; i-- {
		switch remark[i] {
		case '?':
			return true
		case ' ', '\t', '\n', '\r':
			continue
		default:
			break loop
		}
	}
	return false
}

// IsAllUppercase returns true if a sentence has all uppercase characters
func IsAllUppercase(remark string) bool {
	atLeastOneUpperCaseExists := false
	for _, char := range remark {
		if char >= 'a' && char <= 'z' {
			return false
		} else if char >= 'A' && char <= 'Z' {
			atLeastOneUpperCaseExists = true
		}
	}
	return atLeastOneUpperCaseExists
}

// IsSilence returns true if the given remark does not have whitespaces or is empty
func IsSilence(remark string) bool {
	for _, char := range remark {
		switch char {
		case ' ', '\t', '\n', '\r':
			continue
		default:
			return false
		}
	}
	return true
}
