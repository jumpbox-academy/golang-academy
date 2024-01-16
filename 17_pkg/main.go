package main

import (
	"fmt"

	"github.com/jumpbox-academy/golang-academy/animal"
)

func main() {
	a := animal.New()
	fmt.Printf("%T %v\n", a, a)
}
