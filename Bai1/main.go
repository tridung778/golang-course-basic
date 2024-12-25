package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello" + os.Getenv("APP_NAME"))
}
