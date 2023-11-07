package syntax

import (
	"fmt"
)

func TestPanic(a int) {
	deferRecoverPanic(a)
	fmt.Println("print after recover panic")
}

func deferRecoverPanic(a int) {
	defer recoverPanic()
	if a == 0 {
		panic("panic error")
	}
	fmt.Println("print right after panic")
}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("panic recovered")
	}
}
