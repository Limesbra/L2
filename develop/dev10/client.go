package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// структура телнет
type Telnet struct {
	host    string
	port    string
	timeout time.Duration
	conn    net.Conn
}

// Функция Start запускает клиентский сокет для подключения к серверу по протоколу Telnet.
// Она принимает аргументы командной строки --timeout, host и port.
// Если не указаны аргументы, выводится сообщение об ошибке и завершается выполнение программы.
// Возвращает ошибку, возникшую при подключении к серверу или при работе сокета.
func Start() error {
	telnet := &Telnet{}
	telnet.timeout = *flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
		os.Exit(1)
	}

	// Установка значений хоста и порта
	telnet.host = flag.Arg(0)
	telnet.port = flag.Arg(1)

	fmt.Println(telnet)

	// Создание канала для приема сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	dialer := net.Dialer{Timeout: 10 * time.Second}

	// Подключение к серверу
	var err error
	telnet.conn, err = dialer.Dial("tcp", net.JoinHostPort(telnet.host, telnet.port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer telnet.conn.Close()

	// Создание канала для передачи ошибок
	// Запуск горутин SocketWriter и SocketReader
	errChan := make(chan error)
	go telnet.SocketWriter(errChan)
	go telnet.SocketReader(errChan)

	// Ожидание сигнала завершения или возникновения ошибки
	select {
	case <-sigChan:
		return nil
	case err = <-errChan:
		return err
	}

}

// Функция SocketWriter получает ввод с стандартного ввода и отправляет его через сокет.
// Она принимает канал ошибок и возвращает ошибку, возникшую при отправке данных.
func (tel *Telnet) SocketWriter(errChan chan error) {
	for {
		// Создание экземпляра Reader для чтения ввода с стандартного ввода
		inputReader := bufio.NewReader(os.Stdin)
		// Чтение ввода до символа новой строки и сохранение в буфере
		buff, err := inputReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err.Error())
			errChan <- err
			return
		}
		// Отправка буфера через сокет
		if _, err := tel.conn.Write(buff); err != nil {
			errChan <- err
			return
		}
	}

}

// Функция SocketReader получает данные с сервера через сокет и выводит их на стандартный вывод.
// Она принимает канал ошибок и возвращает ошибку, возникшую при чтении данных.
func (tel *Telnet) SocketReader(errChan chan error) {
	for {
		// Создание экземпляра Reader для чтения данных с сервера через сокет
		serverReader := bufio.NewReader(tel.conn)
		for {
			// Чтение данных до символа новой строки и сохранение в буфере
			buff, err := serverReader.ReadBytes('\n')
			if err != nil {
				errChan <- err
				return
			}
			// Вывод данных на стандартный вывод
			fmt.Println(string(buff))
		}
	}

}

func main() {
	err := Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
