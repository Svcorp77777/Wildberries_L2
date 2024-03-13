package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
Реализовать утилиту аналог консольной команды cut (man cut). 
Утилита должна принимать строки через STDIN, разбивать по 
разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "columns" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

type keyCut struct {
	Fields    string
	Delimiter string
	Separated bool
}

func flagArgumentsСonsole() keyCut {
	cutKey := keyCut{}

	flag.StringVar(&cutKey.Fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&cutKey.Delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&cutKey.Separated, "s", false, "только строки с разделителем")

	flag.Parse()

	return cutKey
}

func main() {
	cutKey := flagArgumentsСonsole()

	scann := bufio.NewScanner(os.Stdin)

	for scann.Scan() {
		line := scann.Text()

		if cutKey.Separated && !strings.Contains(line, cutKey.Delimiter) {
			continue
		}

		columns := strings.Split(line, cutKey.Delimiter)
		outputСolumns := selectingColumnsRow(columns, cutKey.Fields)
		fmt.Println(strings.Join(outputСolumns, cutKey.Delimiter))
	}

	if err := scann.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении ввода: %v\n", err)
		os.Exit(1)
	}
}

func selectingColumnsRow(columns []string, columnNumbers string) []string {
	if columnNumbers == "" {
		return columns
	}

	var outputСolumns []string
	selectedСolumns := strings.Split(columnNumbers, ",")

	for _, value := range selectedСolumns {
		i := columnIndex(value, len(columns))
		if i != -1 {
			outputСolumns = append(outputСolumns, columns[i])
		}
	}

	return outputСolumns
}

func columnIndex(column string, maxIndex int) int {
	i := checkPositiv(column)
	if i == 0 || i > maxIndex {
		return -1
	}

	return i - 1
}

func checkPositiv(s string) int {
	num := 0

	fmt.Sscanf(s, "%d", &num)
	if num <= 0 {
		return 0
	}

	return num
}