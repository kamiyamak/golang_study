package main

import (
	"fmt"

	geometry "./tests"
)

type triangle struct {
	size int
}

type coloredTriangle struct {
	triangle
	color string
}

func (t triangle) perimeter() int {
	return t.size * 3
}
func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Size", t.size)
	fmt.Println("Perimeter", t.Perimeter())
}
