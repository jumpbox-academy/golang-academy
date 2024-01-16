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

func (tri *triangle) SetRadius(radius float32) {
	tri.radius = radius
}

func main() {

	tri := triangle{
		radius: 1,
	}
	fmt.Println("Initial Triangle Area", tri.Area())
	tri.SetRadius(3)
	fmt.Println("Assign Radius Triangle Area", tri.Area())
}
