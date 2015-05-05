package main

import "fmt"

func main() {
	addFn := addN(2)
	res := addFn(2)
	fmt.Println(res)
}

func addN(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
