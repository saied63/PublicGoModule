package syntax

import "fmt"

// pointer is a variable address that hold memory address of another variable
// there are two way to decleration of pointer
// firt one is using new keyword

func CreatePointerUsingNew() {
	var p = new(int)
	*p = 10
	fmt.Println("*p : ", *p)
	fmt.Println("&p : ", &p)
}
