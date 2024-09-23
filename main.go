package main

import (
	"fmt"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*Rsvp, 0, 10)
var template = make(map[string]*template.Template, 3)

func loadTemplates() {

	//TODO - load templates here

}

func main() {

	loadTemplates()
	fmt.Println("s")

}
