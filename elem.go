package main

import (
	"fmt"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func ElemHello(name string) *elem.Element {
	return elem.Div(nil,
		elem.Text(fmt.Sprintf("Hello, %s, here are your TODOs", name)),
	)
}

func ElemTodos(todos []Todo) *elem.Element {
	return elem.Ul(nil,
		elem.TransformEach(todos, func(item Todo) elem.Node {
			if item.Done {
				return elem.Li(nil,
					elem.Text(fmt.Sprintf("%s - done", item.Title)))
			} else {
				return elem.Li(nil,
					elem.Text(fmt.Sprintf("%s - nope", item.Title)))
			}
		})...,
	)
}

func ElemGreeting(data Data) *elem.Element {
	return elem.Div(attrs.Props{
		attrs.Class: "greeting",
	},
		ElemHello(data.Hello),
		ElemTodos(data.Todos),
	)
}
