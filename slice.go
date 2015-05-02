package main

import "fmt"

func main() {
	tall := []int{1, 2, 3, 4}
	fmt.Println(tall)

	fmt.Println(tall[2:])
	fmt.Println(tall[1:3])
	fmt.Println(tall[:2])
}
