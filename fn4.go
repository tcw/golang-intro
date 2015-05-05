package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := div(5, 0)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(res)
	}
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Å dele på 0 er tull")
	}
	return a / b, nil
}
