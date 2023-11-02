package syntax

/*
In this tutorial, we will study how to do error handling in golang. As compared to other conventional
languages go doesn’t have exceptions and try-catch for handling the error.
The error handling in golang can be done in two ways

Using type which implements error interface –   it is a conventional way to represent an error and  also idiomatic
Using panic and recover
*/

/* Decleation of error interface
type error interface{
	Error()string
}
*/

import (
	"os"
)

func OpenThisFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	} else {
		defer file.Close()
		return file, nil
	}
}

func OpenAndShowFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	} else {
		defer file.Close()
		return file, nil
	}
}
