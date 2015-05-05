package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	"github.com/unrolled/render"
	"strconv"
)

func init() {
	storage = make(map[int]Ticket)
}

var (
	storage map[int]Ticket
	ren     = render.New(render.Options{
		IndentJSON: true,
	})
)

type Ticket struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	State       string `json:"state"`
}

func (t Ticket) FieldMap() binding.FieldMap {
	return binding.FieldMap{}
}

func main() {
	router := mux.NewRouter().StrictSlash(false)

	posts := router.Path("/ticket").Subrouter()
	posts.Methods("POST").HandlerFunc(PostTicker)
	gets := router.Path("/ticket/{id}").Subrouter()
	gets.Methods("GET").HandlerFunc(GetTicker)
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3001")
}

func PostTicker(w http.ResponseWriter, req *http.Request) {
	ticket := new(Ticket)
	errs := binding.Bind(req, ticket)
	if errs != nil {
		ren.JSON(w, http.StatusMethodNotAllowed, errs)
		return
	}
	storage[ticket.Id] = *ticket
	ren.JSON(w, http.StatusOK, ticket)
}

func GetTicker(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	number, _ := strconv.Atoi(id)
	ren.JSON(w, http.StatusOK, storage[number])
}
