package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	StateNone = iota
	HasSlash
	HasRune
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var lastRune rune

	state := StateNone

	if !utf8.ValidString(s) {
		return "", ErrInvalidString
	}

	for _, currentRune := range s {
		switch state {
		case StateNone:
			switch {
			case currentRune == '\\':
				state = HasSlash
			case unicode.IsDigit(currentRune):
				return "", ErrInvalidString
			default:
				lastRune = currentRune
				state = HasRune
			}

		case HasSlash:
			if unicode.IsDigit(currentRune) || currentRune == '\\' {
				lastRune = currentRune
				state = HasRune
			} else {
				return "", ErrInvalidString
			}

		default:
			if unicode.IsDigit(currentRune) {
				// result.WriteString(strings.Repeat(string(lastRune), int(currentRune-'0'))) // it's better
				times, _ := strconv.Atoi(string(currentRune)) // what would happen? it's a Digit 100%
				result.WriteString(strings.Repeat(string(lastRune), times))

				state = StateNone
			} else {
				result.WriteRune(lastRune)

				if currentRune == '\\' {
					state = HasSlash
				} else {
					lastRune = currentRune
					state = HasRune
				}
			}
		}
	}

	if state == HasRune {
		result.WriteRune(lastRune)

		state = StateNone
	}

	if state != StateNone {
		return "", ErrInvalidString
	}

	return result.String(), nil
}
