package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		c <- "hei"
	}()

	val := <-c
	fmt.Println(val)
}
