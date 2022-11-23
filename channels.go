package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// ch := make(chan int, 100)

	// wg.Add(2)

	// //receive data
	// go func(ch <-chan int) {
	// 	// i := <-ch
	// 	// fmt.Println(i)
	// 	// ch <- 27
	// 	// i = <-ch
	// 	// fmt.Println(i)

	// 	// for i := range ch {
	// 	// 	fmt.Println(i)
	// 	// }

	// 	for {
	// 		if i, ok := <- ch; ok {
	// 			fmt.Println(i)
	// 		} else {
	// 			break
	// 		}
	// 	}

	// 	wg.Done()
	// }(ch)

	// //send data
	// go func(ch chan<- int) {
	// 	i := 12
	// 	ch <- i
	// 	// fmt.Println(<-ch)
	// 	i = 24
	// 	ch <- i

	// 	close(ch)

	// 	// ch <- i
	// 	wg.Done()
	// }(ch)

	// wg.Wait()

	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal2(R2)

	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)
	}
}

func portal1(channel1 chan string) {

	time.Sleep(2 * time.Second)

	channel1 <- "Welcome to channel 1"
}

func portal2(channel2 chan string) {
	time.Sleep(1 * time.Second)

	channel2 <- "Welcome to channel 2"
}
