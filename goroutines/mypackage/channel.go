package mypackage

import (
	"fmt"
	"runtime"
)

func LearnChannel() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}
	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")
	var message1 = <-messages
	fmt.Println(message1)
	var message2 = <-messages
	fmt.Println(message2)
	var message3 = <-messages
	fmt.Println(message3)

	// * buffered channel
	fmt.Println("\nbuffered channel")
	messagesBuffer := make(chan int, 3) // 3 adalah nilai max buffer
	go func() {
		for {
			i := <-messagesBuffer
			fmt.Println("receive data", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messagesBuffer <- i
	}
}
