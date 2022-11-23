package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func worker(id int , wg *sync.WaitGroup)  {
	
	defer wg.Done()

	fmt.Printf("Worker %d starting \n", id)

	time.Sleep(time.Second)
	fmt.Printf("Working %d done \n", id)
}

func main() {
	// go writeMessage()

	//race condition
	// msg := "Hello User"
	// go func() {
	// 	fmt.Println(msg)
	// }()
	// msg = "Hello Deep"

	// msg := "Hello Routine"

	// wg.Add(1)

	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	// msg = "Hello Main"
	// fmt.Println(msg)

	// wg.Wait()

	// time.Sleep(100 * time.Millisecond)


	var wg sync.WaitGroup

	for i := 1 ; i<= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}

// func writeMessage() {
// 	fmt.Println("Hello")
// }
