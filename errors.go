package main

import (
	"errors"
	"fmt")

type error interface {
	Error() string
}

type customError struct {
    Message string
}

func (e customError) Error() string {
	return fmt.Sprintf("Custom Error: %s", e.Message)
}

func main() {
	err := customError{Message: "Something went wrong!"}
	fmt.Println(err.Error())

	var errResponse error= errors.New("This is a standard error")
	fmt.Println(errResponse)
}