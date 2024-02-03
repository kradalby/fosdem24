package main

type Data struct {
	Hello string
	Todos []Todo
}

var Example = Data{
	Hello: "Kristoffer",
	Todos: []Todo{
		{
			Done:  true,
			Title: "Shower",
		},
		{
			Done:  true,
			Title: "Drink (sparkling) water",
		},
		{
			Done:  false,
			Title: "Switch to Yerba Mate",
		},
		{
			Done:  false,
			Title: "Finish the fosdem lightning talk in time",
		},
	},
}
