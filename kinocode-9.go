package main

import (
	"fmt"
)

func main() {
	x := 8
	y := 8
	z := 20

	x++
	y--

	fmt.Println(x)
	fmt.Println(y)

	x += 10
	z += y

	fmt.Println(x)
	fmt.Println(z)

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x >= z)

	fmt.Println(x == z)
	fmt.Println(x != y)

	fmt.Println(x == z && x != y)
	fmt.Println(x != z || y == 1)

}
