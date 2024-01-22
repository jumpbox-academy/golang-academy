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

// Method with Pointer Receiver
func (c *circle) SetRadius(radius float32) {
	c.radius = radius
}

func main() {

	c := circle{
		radius: 1,
	}
	fmt.Println("Initial Triangle Area", c.Area())
	c.SetRadius(3)
	fmt.Println("Assign Radius Triangle Area", c.Area())
}
