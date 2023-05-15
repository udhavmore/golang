package main

import "fmt"

type vehicle interface {
	vname() string
}

type car struct {
	cname string
}

func vname(v vehicle) string {
	return v.vname()
}

func main() {
	c1 := car{"KIA"}
	all_cars := vehicle{c1}
	for cs := range all_cars {
		fmt.Println(vname(cs))
	}
}
