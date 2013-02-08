package main

import (
	"fmt"
	"log"
	"net/http"
)

type Greeting string // HLtype

func (g Greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, g)
}

func main() {
	err := http.ListenAndServe("localhost:4000", Greeting("Hello World"))
	if err != nil {
		log.Fatal(err)
	}
}
