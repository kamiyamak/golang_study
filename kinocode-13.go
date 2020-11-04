package main

import "fmt"

type Student struct {
	name          string
	math, english float64
}

type User struct {
	gender string
	age    int
}

func main() {
	s := Student{english: 70}
	fmt.Println(s)

	u := User{"male", 20}
	fmt.Println(u)
}
