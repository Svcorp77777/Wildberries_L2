package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку.
*/

func main() {

	testCases := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}

	for _, val := range testCases {
		res, err := unpack(val)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("%v -> %v\n", val, res)
	}
}

func unpack(str string) (string, error) {
	var res string

	if len(str) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(str[0])) {
		return "", errors.New("не коректная строка")
	}

	if strings.Contains(str, "\\") {
		res = unpackEscape(str)

		return res, nil
	}

	res = unpackString(str)

	return res, nil
}

func unpackString(str string) string {
	var bild strings.Builder

	for i := 0; i < len(str); i++ {
		if i == len(str)-1 {
			bild.WriteString(string(str[i]))

			break
		}

		if s, ok := strconv.Atoi(string(str[i+1 : i+2])); ok == nil {
			bild.WriteString(strings.Repeat(string(str[i]), s))
			i++

			continue
		}

		bild.WriteString(string(str[i]))
	}

	return bild.String()
}

func unpackEscape(str string) string {
	var bild strings.Builder

	for i := 0; i < len(str); i++ {
		if i == len(str)-1 {
			bild.WriteString(string(str[i]))

			break
		}

		if unicode.IsLetter(rune(str[i])) || unicode.IsDigit(rune(str[i])) {
			bild.WriteString(string(str[i]))

			continue
		}

		if strings.Contains(string(str[i]), "\\") && strings.Contains(string(str[i+1]), "\\") {
			if i < len(str) - 2 && unicode.IsDigit(rune(str[i+2])){
				s, _ := strconv.Atoi(string(str[i+2]))
				bild.WriteString(strings.Repeat(string(str[i+1]), s))

				i += 2

				continue
			} else {
				
				bild.WriteString(string(str[i]))

				i++

				continue
			}
		}

		if strings.Contains(string(str[i]), "\\") && unicode.IsDigit(rune(str[i+1])) {
			if i < len(str)-2 && unicode.IsDigit(rune(str[i+2])) {
				s, _ := strconv.Atoi(string(str[i+2]))
				bild.WriteString(strings.Repeat(string(str[i+1]), s))

				i += 2

				continue
			} else {
				bild.WriteString(string(str[i+1]))

				i++

				continue
			}

			
		}
	}

	return bild.String()
}
