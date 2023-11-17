package syntax

import (
	"fmt"
	"time"
)

func TestChanel() {
	a := make(chan int)

	go SendValueToChan(a)
	fmt.Println("sending...")
	go func() {
		time.Sleep(1 * time.Second)
		go ReciveFromChanfunc(a)
		fmt.Println("recived ...")
	}()

}

func SendValueToChan(ch chan int) {
	ch <- 555
	fmt.Printf("send %d  to chan \n", ch)
}

func ReciveFromChanfunc(ch chan int) {
	recieve := <-ch
	//or
	//_= <-ch
	fmt.Printf("Received %d  fom chan \n", recieve)
}
