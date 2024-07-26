package main

import (
	"fmt"
	"learn-go/arrays"
	"learn-go/conditions"
	"learn-go/slice"
	"learn-go/variables"
)

func main() {
	message := "Hello, World!" // shorthand declaration
	fmt.Println(message)
	variables.Change()
	arrays.Arrays()
	arrays.PriceChange()
	slice.Slice()
	slice.Memory()
	conditions.Nested_if()
}

// Run the code
