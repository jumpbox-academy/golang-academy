package main

import (
	"fmt"
	"strconv"
)

// Scenarios 2
// Animal interface กำหนดเมธอด Speak
type Animal interface {
	Speak() string
}

// Herbivores กำหนดเมธอด Speak สำหรับ Herbivores
type Herbivores struct{}

func (h Herbivores) Speak() string {
	return "I am a Herbivore, I eat plants."
}

// Omnivores กำหนดเมธอด Speak สำหรับ Omnivores
type Omnivores struct{}

func (o Omnivores) Speak() string {
	return "I am an Omnivore, I eat both plants and meat."
}

// Carnivores กำหนดเมธอด Speak สำหรับ Carnivores
type Carnivores struct{}

func (c Carnivores) Speak() string {
	return "I am a Carnivore, I eat meat."
}

// Scenarios 3
type String interface {
	toString() string
}

type nextString string

func (s nextString) toString() string {
	return string(s)
}

type Int int

func (i Int) toString() string {
	return strconv.Itoa(int(i))
}

func main() {

	// Scenario 1
	// declare type interface
	// var testInf interface{}

	// testInf = 6 // assign int
	// fmt.Printf("%T %v\n", testInf, testInf)

	// testInf = "Jumpbox" // assign string
	// fmt.Printf("%T %v\n", testInf, testInf)

	// testInf = true // assign boolean
	// fmt.Printf("%T %v\n", testInf, testInf)

	// testInf = func() string { return "Jumpbox" } // assign function
	// fmt.Printf("%T %v\n", testInf, testInf)

	// Scenarios 2
	// chicken := Herbivores{}
	// dog := Omnivores{}
	// tiger := Carnivores{}

	// fmt.Println("chicken: ", chicken.Speak())
	// fmt.Println("dog: ", dog.Speak())
	// fmt.Println("tiger: ", tiger.Speak())

	// animals := []Animal{chicken, dog, tiger}
	// for _, animal := range animals {
	// 	fmt.Println(animal.Speak())
	// }

	// Scenarios 3
	// Create a toString value
	var s String = nextString("Hello, Go!")
	fmt.Println(s.toString())

	// Create an Int value
	var i String = Int(42)
	fmt.Println(i.toString())

}
