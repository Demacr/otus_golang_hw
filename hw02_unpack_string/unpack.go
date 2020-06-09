package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	fmt.Println("input:", input)

	var i int = 0
	var r = []rune(input)
	var result strings.Builder
	for {
		// Check out of border
		if i == len(r) {
			break
		}
		// Check correctness of new combination
		if unicode.IsLetter(r[i]) {
			// Check if it is a last element or next element is letter
			// (Second check will be only if failed the first part i.e. if next element exists)
			if i+1 == len(r) || unicode.IsLetter(r[i+1]) {
				result.WriteRune(r[i])
				i++
				// Check if the next elemnent is digit
			} else if unicode.IsDigit(r[i+1]) {
				multiply, _ := strconv.Atoi(string(r[i+1]))
				result.WriteString(strings.Repeat(string(r[i]), multiply))
				i = i + 2
			} else {
				return "", ErrInvalidString
			}
		} else {
			return "", ErrInvalidString
		}
	}
	return result.String(), nil
}
