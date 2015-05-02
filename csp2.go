package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("hello")
}

func main() {
	go printHello()
	time.Sleep(time.Second * 2)
}
