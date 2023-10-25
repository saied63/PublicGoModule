package math

import "fmt"

func Divide(a float32, b float32) float32 {
	if b == 0 {
		fmt.Println("secound parameter cant be zero")
		return -1
	}
	return a / b
}
