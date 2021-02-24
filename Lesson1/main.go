package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type MyError struct {
	Err  error
	Time time.Time
}

func main() {
	myErr := MyError{
		Err:  errors.New("new Error"),
		Time: time.Now(),
	}
	fmt.Println(myErr)
	err := openFile("Lesson1/test.txt")
	if err != nil {
		log.Println(err)
	}
	getPanic()
}

func getPanic() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)

		}
	}()

	str := "asdasdas"
	_, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

}

func openFile(fileName string) error {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		return err
	}
	return nil
}
