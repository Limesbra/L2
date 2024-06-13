package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то) +
- pwd - показать путь до текущего каталога +
- echo <args> - вывод аргумента в STDOUT +
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат* +

Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*/

// Функция обертка для командной строки
// слушаем стандартный поток ввода в бесконечном цикле
// при поступлении данных вызываем обработчик  queueCommand
func unixUtil() {
	fmt.Print("Shell:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		if command == "quit" {
			break
		}
		queueCommand(command)

		fmt.Print("enter quit for exit\n")
	}
}

// выполнение каждой команды
func execUtil(c string) (string, error) {
	util := strings.Split(c, " ")
	cmd := exec.Command(util[0], util[1:]...)

	if util[0] == "cd" {
		err := os.Chdir(util[1])
		if err != nil {
			return "", fmt.Errorf("error changing directory: %w", err)
		}
	} else {
		var out strings.Builder
		cmd.Stdout = &out // возвращаем результат в стандартный поток вывода
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("error executing command %s: %w", util[0], err)
		}
		return out.String(), nil
	}

	return "", nil
}

// разбивает команды если это пайплайн и вызывает выполнение каждой команды в отдельности
func queueCommand(c string) {

	cmd := strings.Split(c, "|")
	for _, i := range cmd {
		i = strings.TrimPrefix(i, " ")
		i = strings.TrimSuffix(i, " ")
		out, err := execUtil(i)
		if err != nil {
			log.Fatal("err: ", err)
		}
		fmt.Println(out)
	}
}

func main() {
	unixUtil()
}
