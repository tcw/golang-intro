package main

import "fmt"

type Bil struct {
	reg   string
	dører int
}

func main() {
	bil := Bil{"sp2003456", 4}
	res1 := bil.harBilAntallDører(4)
	res2 := bil.harBilAntallDører(10)
	fmt.Println(res1)
	fmt.Println(res2)
}

func (bil *Bil) harBilAntallDører(dører int) bool {
	return bil.dører == dører
}
