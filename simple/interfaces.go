package main

import "fmt"

// START OMIT
type DuckPerformer interface { // HL
	Quack() string // HL
	Walk() string  // HL
}

func DuckDance(d DuckPerformer) {
	fmt.Printf("%s AND %s\n", d.Quack(), d.Walk())
}

type TyrannosaurusRex int // HL

func (trex *TyrannosaurusRex) Quack() string { // HL
	return "I am quacking"
}

func (trex *TyrannosaurusRex) Walk() string { // HL
	return "I am walking"
}

func main() {
	b := new(TyrannosaurusRex)
	DuckDance(b)
}

// END OMIT
