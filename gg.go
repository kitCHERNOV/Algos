package main

import (
	"fmt"
)
// Add handler for input data
func HandlerInputData(data *int) {
	if *data % 2 == 0{
		fmt.Println("number is even")
	}else {
		fmt.Println("number is odd")
	}
}

func main() {
	fmt.Println("Hello world")
	var c int // my number
	fmt.Scanln(&c)
	fmt.Printf("my number is: %d", c)
}