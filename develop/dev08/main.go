package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
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

*/

func unixUtil() {
	fmt.Print("Shell:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		if command == "quit" {
			break
		}
		out, err := execUtil(command)
		if err != nil {
			log.Fatal("err: ", err)
		}
		fmt.Println(out)
		fmt.Print("enter quit for exit\n")
	}
}

func execUtil(c string) ([]string, error) {
	commands := queueCommand(c)
	fmt.Println(commands)
	var output []string
	var out bytes.Buffer
	for _, item := range commands {
		item = strings.TrimPrefix(item, " ")
		item = strings.TrimSuffix(item, " ")
		util := strings.Split(item, " ")
		// fmt.Println(util, i)
		cmd := exec.Command(util[0], util[1:]...)
		cmd.Stdout = &out
		output = append(output, out.String())
		err := cmd.Run()
		if err != nil {
			fmt.Print(err)
			return nil, fmt.Errorf("error executing command %s: %w", item, err)
		}
	}
	fmt.Println("here", output)

	return output, nil
}

func queueCommand(c string) []string {
	cmd := strings.Split(c, "|")
	return cmd
}

func main() {
	unixUtil()
}
