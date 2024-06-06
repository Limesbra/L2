package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

//Реализовать утилиту wget с возможностью скачивать сайты целиком.

func wget(addres string) error {

	if addres == "" {
		return errors.New("incorrect site name")
	}

	fmt.Println("downloading ...")

	response, err := http.Get(addres)
	if err != nil {
		return errors.New("an error occurred while sending the request")
	}

	defer response.Body.Close()

	file, err := os.Create("index.html")
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return errors.New("error when copying to file")
	}

	fmt.Println("Success!")

	return nil
}

func main() {
	flag.Parse()
	addres := flag.Arg(0)
	fmt.Println(addres)

	err := wget(addres)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
