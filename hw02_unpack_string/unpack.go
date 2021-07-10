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
	prev := ""
	for _, r := range input {
		if unicode.IsDigit(r) {
			if len(prev) == 0 {
				return "", ErrInvalidString
			}

			number, _ := strconv.Atoi(string(r))
			if number == 0 {
				prev = ""
				continue
			}

			builder.WriteString(strings.Repeat(prev, number))
			prev = ""

			continue
		}

		builder.WriteString(prev)
		prev = string(r)
	}

	builder.WriteString(prev)

	return builder.String(), nil
}
