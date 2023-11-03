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
	"fmt"
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

func AddStringToFile(path string, writeTxt string, seek [2]int64) (*os.File, error) {
	var err error
	var file *os.File
	file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	} else {
		defer file.Close()
		_, err = file.Seek(seek[0], int(seek[1]))
		data := make([]byte, 128)
		if err != nil {
			return nil, err
		} else {
			_, err = file.Read(data)
			if err != nil {
				return nil, err
			} else {
				addBytes := []byte(writeTxt)
				data = append(data, addBytes...)
				err = os.WriteFile(path, data, 0644)
				if err != nil {
					return nil, err
				} else {
					fmt.Println("Data Read from file is  : \n", string(data))
					if err = file.Sync(); err != nil {
						return nil, err
					}
					return file, nil
				}
			}
		}
	}
}
