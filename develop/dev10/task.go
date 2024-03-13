package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. 
Если сокет закрывается со стороны сервера, программа должна также завершаться. 
При подключении к несуществующему сервер, программа должна завершаться через timeout

*/
type keyTelnet struct {
	Timeout time.Duration
}

func flagArgumentsСonsole() keyTelnet {
	telnetKey := keyTelnet{}

	timeout := flag.Duration("timeout", 10*time.Second, "Тайм-аут подключения")

	flag.Parse()

	telnetKey.Timeout = *timeout

	return telnetKey
}

func main() {
	telnetKey := flagArgumentsСonsole()

	if flag.NArg() != 2 {
		fmt.Println("Возможно допущена ошибка используйте такой формат: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, telnetKey.Timeout)
	if err != nil {
		fmt.Printf("Возникла ошибка соединения: %s: %v\n", address, err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Подключено к", address)

	go readServer(conn)

	scann := bufio.NewScanner(os.Stdin)
	
	for scann.Scan() {
		line := scann.Text()
		fmt.Fprintf(conn, "%s\n", line)
	}

	if err := scann.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Возникла ошибка при чтении данных: %v\n", err)
		os.Exit(1)
	}
}

func readServer(conn net.Conn) {
	scann := bufio.NewScanner(conn)

	for scann.Scan() {
		line := scann.Text()
		fmt.Println(line)
	}
	
	fmt.Println("Соединение было прервано")
	os.Exit(0)
}