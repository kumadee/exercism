package isbn

import "errors"

// IsValidISBN checks if the given string is a valid ISBN-10 or not.
func IsValidISBN(isbn string) bool {
	sum := 0
	count := 0
	for i := 0; i < len(isbn); i++ {
		if isbn[i] == '-' {
			continue
		}
		digit, error := AtoI(isbn[i])
		if error != nil {
			return false
		}
		if digit == 10 && count != 9 {
			// X is only valid as a check digit
			return false
		}
		sum += digit * (10 - count)
		count++
	}
	return count == 10 && sum%11 == 0
}

// AtoI converts string chars from 0-9 to respective int values and X to 10.
func AtoI(char byte) (int, error) {
	switch {
	case char >= '0' && char <= '9':
		// return int 0 to 9
		return int(char % '0'), nil
	case char == 'X':
		// return 10 for X
		return 10, nil
	default:
		return int(char), errors.New("character is not equal to decimal digit or X")
	}
}
