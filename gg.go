package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	var c int // my number
	fmt.Scanln(&c)
	fmt.Printf("my number is: %d", c)
}