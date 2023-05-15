package main

import "fmt"

type greeter struct {
	name     string
	greeting string
}

func main() {
	g := greeter{
		greeting: "Ola",
		name:     "Peter",
	}
	g.greet()
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
