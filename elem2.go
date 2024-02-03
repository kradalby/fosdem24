package main

import (
	"fmt"

	. "github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func Helloo(name string) *Element {
	return Div(nil,
		Text(fmt.Sprintf("Hello, %s, here are your TODOs", name)),
	)
}

func Todoss(todos []Todo) *Element {
	return Ul(nil,
		TransformEach(todos, func(item Todo) Node {
			if item.Done {
				return Li(nil,
					Text(fmt.Sprintf("%s - done", item.Title)))
			} else {
				return Li(nil,
					Text(fmt.Sprintf("%s - nope", item.Title)))
			}
		})...,
	)
}

func Greetingg(data Data) *Element {
	return Div(attrs.Props{
		attrs.Class: "greeting",
	},
		Helloo(data.Hello),
		Todoss(data.Todos),
	)
}
