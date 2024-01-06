package main

import (
	"fmt"
)

func main() {

	// Declare variable with non initial a value
	// var number int
	// var text string
	// var isBool bool

	// Declare variable with initial a value
	var number int = 6
	var text string = "Jumpbox"
	var isBool bool = true

	fmt.Println("number:", number, "text:", text, "bool:", isBool)
	fmt.Printf("number: %d text: %s bool: %t\n", number, text, isBool)

	var stringReturn = fmt.Sprint("number:", number, "text:", text, "bool:", isBool)
	var stringReturnFormat = fmt.Sprintf("number: %d text: %s bool: %t", number, text, isBool)

	fmt.Println("stringReturn:", stringReturn)
	fmt.Println("stringReturnFormat:", stringReturnFormat)
}
