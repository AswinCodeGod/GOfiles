package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"

	// currentMessage := <-channel
	// fmt.Println("this is the current message: ", currentMessage)

	// if currentMessage == "pong" {
	// 	fmt.Println("Foo has received: pong")
	// 	bar(channel)
	// }

	message := "ping"

	fmt.Println("Foo is sending: ", message)
	channel <- message

}

func bar(channel chan string) {
	// TODO: Write an infinite loop of receiving "pings" and sending "pongs"

	// currentMessage := <-channel
	// fmt.Println("this is the current message: ", currentMessage)

	// if currentMessage == "ping" {
	// 	fmt.Println("Foo has received: ping")
	// 	bar(channel)
	// }

	message := "pong"

	fmt.Println("Bar is sending: ", message)
	channel <- message

}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	message := make(chan string)

	go foo(message)
	go bar(message)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
