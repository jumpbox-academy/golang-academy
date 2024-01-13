package main

import "fmt"

func main() {
	fmt.Println("Hello World golang")
}

func noReturn() {
	fmt.Println("No Return")
}

func returnString() string {
	return "Jumpbox"
}

// return single value
func add(a, b int) int {
	return a + b
}

// return multiple value
func swap(a, b int) (int, int) {
	return b, a
}
