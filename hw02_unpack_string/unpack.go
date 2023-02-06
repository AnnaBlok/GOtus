package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

var ErrInvalidEscaping = errors.New("invalid escaping")

func Unpack(s string) (string, error) {
	chs := []rune(s)
	var unpacked strings.Builder

	for i := 0; i < len(chs); i++ {
		switch {
		case unicode.IsDigit(chs[i]) && !isEscaped(i, chs):
			if isNumPositionIncorrect(i, chs) {
				return "", ErrInvalidString
			}
		case string(chs[i]) == `\` && !isEscaped(i, chs):
			if isSlashPositionIncorrect(i, chs) {
				return "", ErrInvalidEscaping
			}
		default:
			unpacked.WriteString(strings.Repeat(string(chs[i]), getNumberOfChToRepeat(i, chs)))
		}
	}
	return unpacked.String(), nil
}

func isEscaped(i int, chs []rune) bool {
	var count int
	if i > 0 {
		for ; i >= 1 && string(chs[i-1]) == `\`; i-- {
			count++
		}
		if (count % 2) == 1 {
			return true
		}
	}
	return false
}

func isNumPositionIncorrect(i int, chs []rune) bool {
	if i == 0 || (unicode.IsDigit(chs[i-1]) && !isEscaped(i-1, chs)) {
		return true
	}
	return false
}

func isSlashPositionIncorrect(i int, chs []rune) bool {
	if (i == (len(chs)-1) && !isEscaped(i, chs)) || (!unicode.IsDigit(chs[i+1]) && string(chs[i+1]) != `\`) {
		return true
	}
	return false
}

func getNumberOfChToRepeat(i int, chs []rune) int {
	var count int
	if i < (len(chs) - 1) {
		if unicode.IsDigit(chs[i+1]) {
			count, _ = strconv.Atoi(string(chs[i+1]))
		} else {
			count = 1
		}
	} else {
		count = 1
	}
	return count
}
