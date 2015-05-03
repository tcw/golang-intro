package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	time.Sleep(time.Second * 5)
}
