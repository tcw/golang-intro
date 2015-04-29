package main

import "fmt"

var (
	tekst string = "hello"
)

func main() {
	fmt.Printf("tekst(main) : %s\n", tekst)
	fn1()
}

func fn1() {
	fmt.Printf("tekst(fn1) : %s", tekst)
}
