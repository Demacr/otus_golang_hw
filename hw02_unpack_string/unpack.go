package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var i int = 0
	var r = []rune(input)
	var result strings.Builder
	for {
		// Check out of border
		if i == len(r) {
			break
		}
		// Check correctness of first element of new combination
		if unicode.IsLetter(r[i]) {
			// Check if it is a last element
			if i+1 == len(r) {
				result.WriteRune(r[i])
				i++
				continue
			}
			// Else current index not last, so let's check what's next
			switch next := r[i+1]; {
			case unicode.IsLetter(next):
				result.WriteRune(r[i])
				i++
			case unicode.IsDigit(next):
				multiply, _ := strconv.Atoi(string(next))
				result.WriteString(strings.Repeat(string(r[i]), multiply))
				i += 2
			default:
				return "", ErrInvalidString
			}
		} else {
			return "", ErrInvalidString
		}
	}
	return result.String(), nil
}
