package main

import "fmt"

func main() {
	res := add(2, 3)
	fmt.Println(res)
}

func add(a, b int) (c int) {
	c = a + b
	return
}
