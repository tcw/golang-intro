package main

import (
	"fmt"
	"math"
)

func main() {
	c := applyFn(-40, math.Abs)
	fmt.Println(c)
}

func applyFn(a float64, f func(float64) float64) float64 {
	return f(a)
}
