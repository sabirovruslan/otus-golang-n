package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var builder strings.Builder
	var prev strings.Builder
	for _, r := range input {
		if unicode.IsDigit(r) {
			if prev.Len() == 0 {
				return "", ErrInvalidString
			}

			number, _ := strconv.Atoi(string(r))
			builder.WriteString(strings.Repeat(prev.String(), number))
			prev.Reset()
			continue
		}

		builder.WriteString(prev.String())
		prev.Reset()
		prev.WriteString(string(r))
	}

	builder.WriteString(prev.String())

	return builder.String(), nil
}
