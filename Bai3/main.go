package main

import (
	"fmt"
	"math"
)

func loop() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	// loop()
	// fmt.Println(sqrt(2), sqrt(-4))

}
