package main

import "fmt"

var (
	tekst string = "hello"
)

func main() {
	fmt.Printf("tekst-main : %s\n", tekst)
	printTekst()
}

func printTekst() {
	fmt.Printf("tekst-printTekst : %s", tekst)
}
