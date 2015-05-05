package main

import "fmt"

type Bil struct {
	reg   string
	dører int
}

func main() {
	bil := Bil{"sp2003456", 4}
	res1 := bil.harDører(4)
	fmt.Println(res1)
}

func (bil *Bil) harDører(dører int) bool {
	return bil.dører == dører
}
