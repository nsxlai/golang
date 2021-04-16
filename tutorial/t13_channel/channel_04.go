package main

import (
	"fmt"
	"time"
)

func main() {
	// example 1
	fmt.Println("Example 1")
	messages1 := make(chan string)

	go func() { messages1 <- "ping" }()

	msg := <-messages1
	fmt.Println(msg)

	// example 2: channel buffering
	fmt.Println("Example 2")
	messages2 := make(chan string, 2)

	messages2 <- "buffered"
	messages2 <- "channel"

	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	// example 3: channel synchronization
	fmt.Println("Example 3")
	done := make(chan bool, 1)
	go worker(done)

	<-done

	// example 4: channel direction
	fmt.Println("Example 4")
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)

}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}