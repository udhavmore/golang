package main

import "fmt"

func main() {
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	// complex structs are dereferenced automatically
	ms.foo = 420
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
