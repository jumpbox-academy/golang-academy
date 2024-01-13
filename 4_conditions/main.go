package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Short statement
	if n, err := strconv.Atoi("5s"); err != nil { // in-case error
		fmt.Println("error short statement with err:", err)
	} else {
		fmt.Println("else short statement with n:", n)
	}

	// Full statement
	n, err := strconv.Atoi("5") // in-case success
	if err != nil {
		fmt.Println("error full statement with err:", err)
	} else {
		fmt.Println("else full statement with n:", n)
	}
}
