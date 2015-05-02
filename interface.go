package main

import "fmt"

type Figur interface {
	Areal() int
}

type Rektangel struct {
	lengde, bredde int
}

type Kvaderat struct {
	side int
}

func (r Rektangel) Areal() int {
	return r.lengde * r.bredde
}

func (k Kvaderat) Areal() int {
	return k.side * k.side
}

func main() {
	r := Rektangel{5, 3}
	k := Kvaderat{5}

	var f Figur
	f = r
	fmt.Println("Areal av Figur(Rektangel): ", f.Areal())
	f = k
	fmt.Println("Areal av Figur(Kvaderat): ", f.Areal())
}
