package main

import (
	"fmt"
)

func main() {
	a := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tanaka"}}

	fmt.Println(a[0][0])
	fmt.Println(a[0][1])
	fmt.Println(a[1][0])
	fmt.Println(a[1][1])
}
