package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд
Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена
команда выхода (например \quit).
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		if !scanner.Scan() {
			break
		}

		str := scanner.Text()
		args := strings.Fields(str)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "pwd":
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Не удалось получить путь до текущего каталога:", err)
			}
			fmt.Println(cwd)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "cd":
			if len(args) < 2 {
				fmt.Println("cd необходимо использовать в следующем формате: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Не удалось сменить каталог:", err)
			}
		case "ps":
			out, err := exec.Command("ps").Output()
			if err != nil {
				fmt.Println("Не удалось выполнить команду ps:", err)
			}
			fmt.Println(string(out))
		case "kill":
			if len(args) < 2 {
				fmt.Println("kill необходимо использовать в следующем формате: kill <pid>")
				continue
			}
			pid := args[1]
			err := exec.Command("kill", pid).Run()
			if err != nil {
				fmt.Println("Не удалось прервать процесс:", err)
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("Не удалось выполнить команду:", err)
			}
		}
	}
}
