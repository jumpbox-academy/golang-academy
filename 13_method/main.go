package main

import "fmt"

const PI = 3.14

type circle struct {
	radius float32
}

// Method
func (c circle) Area() float32 {
	return PI * c.radius * c.radius
}

func (c circle) Circumference() float32 {
	return 2 * PI * c.radius
}

func main() {

	c := circle{
		radius: 3,
	}
	fmt.Println("Triangle Area", c.Area())
	fmt.Println("Triangle Circumference", c.Circumference())
}
