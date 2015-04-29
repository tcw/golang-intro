package main

import (
	"fmt"
	"reflect"
)

func main() {
	tekst := "hello"
	tall := 42
	desiamlTall := 4.2

	fmt.Printf("%s -> %s\n", reflect.TypeOf(tekst), tekst)
	fmt.Printf("%s -> %d\n", reflect.TypeOf(tall), tall)
	fmt.Printf("%s -> %f\n", reflect.TypeOf(desiamlTall), desiamlTall)
}
