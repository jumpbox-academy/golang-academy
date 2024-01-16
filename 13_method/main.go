package main

import "fmt"

const PI = 3.14

type triangle struct {
	radius float32
}

// Method
func (tri triangle) Area() float32 {
	return PI * tri.radius * tri.radius
}

func main() {

	tri := triangle{
		radius: 3,
	}
	fmt.Println("Triangle Area", tri.Area())
}
