package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("Starting to panic")
	// Anonymous function
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic("System panicked again!!!")
		}
	}()
	panic("system panicked!!!")
	fmt.Println("Done panicking")
}
