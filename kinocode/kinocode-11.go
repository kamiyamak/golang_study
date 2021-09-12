package main

import "fmt"

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	sum := 0

	for i := 0; i <= 4; i++ {
		sum += arr[i]
	}
	fmt.Println(sum)

	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	k := 1
	for {
		fmt.Println(k)
		if k == 10 {
			break
		}
		k++
	}

	for h := 1; h <= 10; h++ {
		fmt.Println(h)
	}
}
