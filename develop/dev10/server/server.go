/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout

telnet 23 tcp port*/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//слушаем порт 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	fmt.Println("Server started. Listening on port 8080...")

	for {
		// ожидаем подключение.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// создаемм горутину
		go func(c net.Conn) {
			// дублируем данные.
			io.Copy(c, c)
			c.Close()
			// fmt.Println("Connection closed")
		}(conn)
	}
}
