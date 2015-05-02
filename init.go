package main

import "fmt"

func main() {
	fmt.Println("deretter Main")
}

func init() {
	fmt.Print("Init kjører først, ")
}
