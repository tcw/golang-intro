package main

import "fmt"

func main() {
	res := add(2, 3)
	fmt.Println(res)

	resNamed := addNamed(2, 3)
	fmt.Println(resNamed)
}

func add(a, b int) int {
	c := a + b
	return c
}

func addNamed(a, b int) (c int) {
	c = a + b
	return
}