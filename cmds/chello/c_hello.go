package main

import (
	"C"
	"fmt"
)

// we require one but can be empty
func main() {}

//export Hello
func Hello() {
	fmt.Println("Hello, world! from C")
}
