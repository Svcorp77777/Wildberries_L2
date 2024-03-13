package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort
(man sort — смотрим описание и основные параметры): на входе подается
файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в
качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

type keySort struct {
	Column         int
	NumericValue   bool
	Reverse        bool
	DuplicateLines bool
	MonthName      bool
	IgnoreSpaces   bool
	CheckSort      bool
	SortNumSuff    bool
}

type sortFile struct {
	lines      []string
	parKeySort keySort
}

func parametersKeySort() keySort {
	sortKey := keySort{}

	flag.IntVar(&sortKey.Column, "k", 0, "сортировка по колонки по умолчанию 0")
	flag.BoolVar(&sortKey.NumericValue, "n", false, "сортировка по числовому значению")
	flag.BoolVar(&sortKey.Reverse, "r", false, "сортировка в обратном порядке")
	flag.BoolVar(&sortKey.DuplicateLines, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&sortKey.MonthName, "M", false, "сортировка по названию месяца")
	flag.BoolVar(&sortKey.IgnoreSpaces, "b", false, "игнорирование хвостовых пробелов")
	flag.BoolVar(&sortKey.CheckSort, "c", false, "проверка отсортированы ли данные")
	flag.BoolVar(&sortKey.SortNumSuff, "h", false, "сортировка по числовому значению с учетом суффиксов")

	flag.Parse()

	return sortKey
}

func main() {
	inputFile := flag.String("file", "", "путь к входному файлу")
	parKeySort := parametersKeySort()

	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Не указан путь к входному файлу")

		return
	}

	lines, err := readingDataFromFile(*inputFile)
	if err != nil {
		fmt.Println("Не удалось прочитать файл:", err)

		return
	}

	copyLines := make([]string, len(lines))
	copy(copyLines, lines)

	if parKeySort.DuplicateLines {
		lines = deleteDuplicate(lines)
	}

	sort.Sort(sortFile{lines, parKeySort})

	if parKeySort.CheckSort && sortData(copyLines, parKeySort) {
		fmt.Println("Данные отсортированны")

		return
	} else if parKeySort.CheckSort && !sortData(copyLines, parKeySort) {
		fmt.Println("Данные не отсортированны")

		return
	}

	outputFile := "task_" + *inputFile
	err = writingDataToFile(outputFile, lines)
	if err != nil {
		fmt.Println("Ошибка при записи файла: ", err)
		return
	}

	fmt.Println("Завершено успешно. Данные записаны в файл:", outputFile)
}

func readingDataFromFile(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	return lines, scan.Err()
}

func deleteDuplicate(lines []string) []string {
	see := make(map[string]struct{})
	res := make([]string, 0, len(lines))

	for _, line := range lines {
		if _, exists := see[line]; !exists {
			res = append(res, line)
			see[line] = struct{}{}
		}
	}

	return res
}

func sortData(lines []string, parKeySort keySort) bool {
	sort := sortFile{lines, parKeySort}

	for i := 1; i < len(lines); i++ {
		if sort.Less(i, i-1) {
			return false
		}
	}

	return true
}

func writingDataToFile(inputFile string, lines []string) error {
	file, err := os.Create(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func (s sortFile) Len() int {
	return len(s.lines)
}

func (s sortFile) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

func (s sortFile) Less(i, j int) bool {
	line1 := s.lines[i]
	line2 := s.lines[j]

	if s.parKeySort.NumericValue {
		num1, err1 := extractNumericValue(line1, s.parKeySort.Column, s.parKeySort.SortNumSuff)
		num2, err2 := extractNumericValue(line2, s.parKeySort.Column, s.parKeySort.SortNumSuff)

		if err1 == nil && err2 == nil {
			if num1 < num2 {
				return !s.parKeySort.Reverse
			} else if num1 > num2 {
				return s.parKeySort.Reverse
			}
		}
	}

	if s.parKeySort.MonthName {
		month1, err1 := parseMonthName(line1, s.parKeySort.Column)
		month2, err2 := parseMonthName(line2, s.parKeySort.Column)

		if err1 == nil && err2 == nil {
			if month1 < month2 {
				return !s.parKeySort.Reverse
			} else if month1 > month2 {
				return s.parKeySort.Reverse
			}
		}
	}

	if s.parKeySort.IgnoreSpaces {
		line1 = strings.TrimSpace(line1)
		line2 = strings.TrimSpace(line2)
	}

	if !s.parKeySort.Reverse {
		return line1 < line2
	} else {
		return line1 > line2
	}
}

func extractNumericValue(line string, columnIndex int, numericSuffix bool) (float64, error) {
	fields := strings.Fields(line)

	if columnIndex >= len(fields) {
		return 0, fmt.Errorf("столбца с таким индексом не существует")
	}

	value := fields[columnIndex]
	if numericSuffix {
		return parNumValueWithSuff(value)
	}

	return strconv.ParseFloat(value, 64)
}

func parNumValueWithSuff(value string) (float64, error) {
	suffixes := map[string]float64{
		"K": 1e3,
		"M": 1e6,
		"G": 1e9,
		"T": 1e12,
	}

	for suffix, multiplier := range suffixes {
		if strings.HasSuffix(value, suffix) {
			numStr := strings.TrimSuffix(value, suffix)
			num, err := strconv.ParseFloat(numStr, 64)

			if err != nil {
				return 0, err
			}

			return num * multiplier, nil
		}
	}

	return strconv.ParseFloat(value, 64)
}

func parseMonthName(line string, columnIndex int) (time.Month, error) {
	fields := strings.Fields(line)
	if columnIndex >= len(fields) {
		return 0, fmt.Errorf("индекс столбца вне диапазона")
	}

	monthStr := strings.ToLower(fields[columnIndex])
	switch monthStr {
	case "январь":
		return time.January, nil
	case "февраль":
		return time.February, nil
	case "март":
		return time.March, nil
	case "апрель":
		return time.April, nil
	case "май":
		return time.May, nil
	case "июнь":
		return time.June, nil
	case "июль":
		return time.July, nil
	case "август":
		return time.August, nil
	case "сентябрь":
		return time.September, nil
	case "октябрь":
		return time.October, nil
	case "ноябрь":
		return time.November, nil
	case "декабрь":
		return time.December, nil
	default:
		return 0, fmt.Errorf("неверный месяц: %s", monthStr)
	}
}
