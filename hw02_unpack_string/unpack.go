package hw02unpackstring

import (
	"errors"
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
	if !utf8.ValidString(s) {
		return "", ErrInvalidString
	}

	var result strings.Builder

	lastRune := rune(0)
	state := StateNone

	for _, currentRune := range s {
		switch state {
		case StateNone:
			if currentRune == '\\' {
				state = HasSlash
			} else if unicode.IsDigit(currentRune) {
				return "", ErrInvalidString
			} else {
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
				result.WriteString(strings.Repeat(string(lastRune), int(currentRune-'0')))

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
