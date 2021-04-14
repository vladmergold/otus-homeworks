package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here

	// Пустая строка
	if len(str) == 0 {
		return "", nil
	}

	str += "\n"         // для дальнейшего удобного перебора, чтобы проверить все элементы
	rStr := []rune(str) // преобразовали в слайс

	// Невалидная строка (начинается с цифры)
	if unicode.IsDigit(rStr[0]) {
		return "-1", ErrInvalidString
	}

	var bOut strings.Builder
	var bSp strings.Builder

	var slash bool
	// a5bc\4\5x abc\45x abc\\\5\x `abc\n4`
	for i := range rStr {
		var iCount int
		if string(rStr[i]) == `\` && !slash {
			bOut.WriteString(strings.Repeat(bSp.String(), 1))
			bSp.Reset()
			slash = true
			continue
		}

		if slash {
			if string(rStr[i]) != `\` && !unicode.IsDigit(rStr[i]) {
				return "-2", ErrInvalidString // косяк
			}
			bSp.WriteRune(rStr[i])
			slash = false
			continue
		}

		if unicode.IsDigit(rStr[i]) {
			if bSp.Cap() == 0 {
				return "-3", ErrInvalidString
			}
			iCount, _ = strconv.Atoi(string(rStr[i])) // считаем
			bOut.WriteString(strings.Repeat(bSp.String(), iCount))
			bSp.Reset()
		} else {
			bOut.WriteString(strings.Repeat(bSp.String(), 1))
			bSp.Reset()
			bSp.WriteRune(rStr[i])
		}
	}
	return bOut.String(), nil
}
