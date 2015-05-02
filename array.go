package main

import "fmt"

func main() {
	var a [4]int
	a[0] = 1
	fmt.Println(a)

	b := [...]int{1, 2}
	fmt.Println(b)
}
