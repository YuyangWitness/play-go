package main

import (
	"fmt"
	"time"
)

func main() {
	doneChan := make(chan string)

	go func() {
		doneChan <- "message"
	}()

	<-doneChan
}

func test1() {
	theMine := []string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	go func(theMine []string) {
		for _, m := range theMine {
			oreChan <- m
			fmt.Println("send", m)
		}
	}(theMine)

	go func() {
		for i := 0; i < 3; i++ {
			received := <-oreChan
			fmt.Println("received", received)
		}
	}()

	<-time.After(time.Second * 5) // Again, ignore this for now
}
