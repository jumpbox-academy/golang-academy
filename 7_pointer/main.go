package main

import "fmt"

func main() {
	var s string
	var p *string
	p = &s
	*p = "Jumpbox"

	fmt.Println("Pointer Adress: ", p)
	fmt.Println("Pointer Jump to value: ", *p)

}
