package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 12

	if age := x + y; age >= 20 {
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}

	// test
	a := 7

	if a >= 5 && a < 10 {
		fmt.Println("match")
	} else {
		fmt.Println("unmatch")
	}
}
