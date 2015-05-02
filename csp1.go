package main

import "fmt"

func printHello() {
	fmt.Println("hello")
}

func main() {
	go printHello()
}
