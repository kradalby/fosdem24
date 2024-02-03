package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/flosch/pongo2/v6"
)

type Todo struct {
	Done  bool
	Title string
}

func main() {
	// Go template

	buf := new(bytes.Buffer)
	gotemp := template.Must(template.ParseFiles("hello.gotemp.html"))
	err := gotemp.Execute(buf, Example)
	if err != nil {
		log.Fatalf("failed to render gotemplate: %s", err)
	}

	fmt.Println("gotemplate:")
	fmt.Println(buf.String())

	// pongo2
	buf = new(bytes.Buffer)
	pongotemp := pongo2.Must(pongo2.FromFile("hello.pongo.html"))
	pongotemp.ExecuteWriter(pongo2.Context{"data": Example}, buf)

	fmt.Println("pango2:")
	fmt.Println(buf.String())

	// templ
	buf = new(bytes.Buffer)
	templComponenet := Greeting(Example)
	err = templComponenet.Render(context.TODO(), buf)
	if err != nil {
		log.Fatalf("failed to render gotemplate: %s", err)
	}

	fmt.Println("templ:")
	fmt.Println(buf.String())

	// elem-go
	elem := ElemGreeting(Example)

	fmt.Println("elem-go-1:")
	fmt.Println(elem.Render())

	elem2 := Greetingg(Example)

	fmt.Println("elem-go-2:")
	fmt.Println(elem2.Render())
}
