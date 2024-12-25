package main

import "fmt"

func main() {
	a := 10
	var b bool = true
	const c = 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(add(42, 13))
}

func add(x, y int) int {
	return x + y
}
