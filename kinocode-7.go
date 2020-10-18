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

	var stringA string = "Hello,World!"
	stringB := "Hello,World!"
	fmt.Println(stringA)
	fmt.Println(reflect.TypeOf(stringA))
	fmt.Println(stringB)
	fmt.Println(reflect.TypeOf(stringB))

	a := 10
	b := 1
	numbool := a < b

	fmt.Println(numbool)
	fmt.Println(reflect.TypeOf(numbool))
}
