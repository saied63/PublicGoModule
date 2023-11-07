package main

func TestPanic(a int) {
	if a == 0 {
		panic("panic error")
	}
}
