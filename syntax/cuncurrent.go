package syntax

import (
	"fmt"
	"time"
)

func TestGoRoutine() {
	fmt.Println("befor call goroutin 1")
	///////////////////////
	go callGoRoutine1()
	// we can sleep some times for execute callGoRoutine1
	time.Sleep(1 * time.Second)
	///////////////////////
	fmt.Println("after call goRoutin 1")
}

func callGoRoutine1() {
	fmt.Println("inside goRoutine 1")
}

func CreateMultipleGoRoutine() {
	for i := 0; i < 100; i++ {
		go execute(i)
	}
}

func execute(a int) {
	fmt.Println(a)
}
