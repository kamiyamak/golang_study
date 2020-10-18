package main

import (
	"fmt"
	"reflect"
)

func main() {
	num2 := 2
	var num02 int = 1234567890
	num03 := 1.23
	var num04 float64 = 1.23456789

	fmt.Println(reflect.TypeOf(num2))
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(reflect.TypeOf(num03))
	fmt.Println(reflect.TypeOf(num04))

	var string_a string = "Hello,World!"
	string_b := "Hello,World!"
	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))

	a := 10
	b := 1
	num_bool := a < b

	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))
}
