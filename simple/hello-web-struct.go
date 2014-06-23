package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT

// package and import statements here...

type Greeting struct {
	Message string
	Count   int
}

func (g *Greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.Count++
	m := fmt.Sprintf("%s. Fired %d times.", g.Message, g.Count) // or var m string = ...
	//fmt.Printf("Path: %s\n", r.URL.Path)
	fmt.Fprint(w, m)
}
func main() {
	var g *Greeting = &Greeting{Message: "Hello World"} // or new(Greeting)...
	err := http.ListenAndServe("localhost:4000", g)
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT
