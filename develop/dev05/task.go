package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой 
(man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

type keyGrep struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Pattern    string
	FilePaths  []string
}

func flagArgumentsСonsole() keyGrep {
	grepKey := keyGrep{}

	flag.IntVar(&grepKey.After, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&grepKey.Before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&grepKey.Context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&grepKey.Count, "c", false, "количество строк")
	flag.BoolVar(&grepKey.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&grepKey.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&grepKey.Fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&grepKey.LineNum, "n", false, "печатать номера строк")

	flag.Parse()

	grepKey.FilePaths = flag.Args()

	if len(grepKey.FilePaths) > 0 {
		grepKey.Pattern = grepKey.FilePaths[0]
		grepKey.FilePaths = grepKey.FilePaths[1:]
	}

	return grepKey
}

func main() {
	grepKey := flagArgumentsСonsole()

	if len(grepKey.FilePaths) == 0 {
		fmt.Println("Укажите путь к входному файлу.")

		return
	}

	for _, filePath := range grepKey.FilePaths {
		lines, err := readingDataFromFile(filePath)

		if err != nil {
			fmt.Printf("Не удалось прочитать файл %s: %v\n", filePath, err)

			return
		}

		existingLine := grepFilter(lines, grepKey)
		printResult(existingLine, grepKey)
	}
}

func readingDataFromFile(filePath string) ([]string, error) {
	var lines []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	return lines, scan.Err()
}

func grepFilter(lines []string, grepKey keyGrep) []string {
	var result []string
	pattern := grepKey.Pattern

	if grepKey.IgnoreCase {
		pattern = "(?i)" + pattern
	}

	if grepKey.Fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Ошибка в регулярном выражении: %v\n", err)

		return result
	}

	for i, line := range lines {
		matched := re.MatchString(line)

		if grepKey.Invert {
			matched = !matched
		}

		if matched {
			result = append(result, line)

			if grepKey.After > 0 && i+grepKey.After < len(lines) {
				result = append(result, lines[i+1:i+1+grepKey.After]...)
			}

			if grepKey.Before > 0 && i-grepKey.Before >= 0 {
				result = append(result, lines[i-grepKey.Before:i]...)
			}

			if grepKey.Context > 0 && i-grepKey.Context >= 0 && i+grepKey.Context < len(lines) {
				result = append(result, lines[i-grepKey.Context:i]...)
				result = append(result, lines[i+1:i+1+grepKey.Context]...)
			}
		}
	}

	return result
}

func printResult(lines []string, grepKey keyGrep) {
	if grepKey.Count {
		fmt.Printf("Количество совпадений := %d\n", len(lines))
	} else {
		for i, line := range lines {
			if grepKey.LineNum {
				fmt.Printf("%d: ", i+1)
			}

			fmt.Println(line)
		}
	}
}