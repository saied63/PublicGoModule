package syntax

import (
	"fmt"
)

func PrintThisFromCuncurrent(s string) {
	fmt.Println("i am printing this from concurrent : " + s)
}
func AddAnotherToConcurrent(s string) {
	fmt.Println("i am secound print from concurrent")
}
