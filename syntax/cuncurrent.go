package syntax

import (
	"fmt"
	"runtime"
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
	for i := 0; i < 5; i++ {
		go execute(i)
	}
}

func execute(a int) {
	fmt.Println(a)
}

func PrintNumberOfVirtualCpusTHreads() {
	fmt.Println("/************** total virtual threads is : ")
	fmt.Println(runtime.NumCPU())
	fmt.Println("**************/")
	// call an anymous function
	go func() {
		fmt.Println("anumous func .... ")
	}()
}
