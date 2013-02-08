package main

import (
	"fmt"
	"log"
	"net/http"
)
type Greeting struct {
	Message string
	Count   int
}
func (g *Greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.Count++
	m := fmt.Sprintf("%s. Fired %d times.", g.Message, g.Count) // or var m string = ...
	//fmt.Printf("Full url: %s\n", r.URL)
	fmt.Fprint(w, m)
}
func main() {
	var g *Greeting = &Greeting{Message: "Hello World"} // or new(Greeting)...
	err := http.ListenAndServe("localhost:4000", g)
	if err != nil {
		log.Fatal(err)
	}
}
