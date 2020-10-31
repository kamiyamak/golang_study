package main

import "fmt"

func sayHello(greeting string) {
	fmt.Println(greeting)
}

func cal(x, y int) (a int, b int) {
	a = x + y
	b = x * y
	return
}

func main() {

	func(greeting string) {
		fmt.Println(greeting)
	}("Good evening")

	result01, result02 := cal(10, 5)
	fmt.Println(result01, result02)
}
