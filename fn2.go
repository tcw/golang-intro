package main

import "fmt"

func main() {
	res := addN(2)
	c := addtoFn(40, res)
	fmt.Println(c)
}

func addN(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func addtoFn(a int, f func(int) int) int {
	return f(a)
}
