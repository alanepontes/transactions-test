package main

import (
	"fmt"
)

type CustomError struct {
	argument  string
	text string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s - %s", e.argument, e.text)
}