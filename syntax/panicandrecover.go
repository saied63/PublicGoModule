package syntax

import (
	"fmt"
)

func TestPanic(a int) {
	if a == 0 {
		panic("panic error")
	}
	fmt.Println("print after panic")
}
