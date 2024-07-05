package main

import (
	"fmt"
	"learn-go/variables"
	"learn-go/arrays"
	"learn-go/slice"
)

func main() {
	message := "Hello, World!" // shorthand declaration
	fmt.Println(message)
	variables.Change()
	arrays.Arrays()
	arrays.PriceChange()
	slice.Slice()
	slice.Memory()
}

// Run the code
