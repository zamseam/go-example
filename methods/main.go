package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width = 11
	return r.width * r.height
}

func (r rect) perim() int {
	r.width = 15
	return 2*r.width + 2*r.height
}

func main() {
	c := rect{width:10, height: 5}
	// fmt.Println("area: ", c.area())
	// fmt.Println(c.width)
	fmt.Println("perim:", c.perim())
	fmt.Println(c.width)

	// rp := &c
	// fmt.Println("area: ", rp.area())
	// fmt.Println("perim:", rp.perim())
}