package main

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"strconv"
)

func init() {
	storage = make(map[int]Ticket)
}

var (
	storage map[int]Ticket
)

type Ticket struct {
	Number      int    `json:"number"`
	Description string `json:"description"`
	State       string `json:"state"`
}

func ReadStatus(p martini.Params) Ticket {

	ticket := Ticket{
		Number:      645,
		Description: "A dummy customer ticket " + p["id"],
		State:       "resolved",
	}
	return ticket
}

func main() {

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/status/:id", func(r render.Render, params martini.Params) {
		id, _ := strconv.Atoi(params["id"])
		r.JSON(200, storage[id])
	})
	m.Post("/status", binding.Bind(Ticket{}), func(ticket Ticket, r render.Render) {
		storage[ticket.Number] = ticket
		r.JSON(200, ticket)
	})
	m.Run()
}
