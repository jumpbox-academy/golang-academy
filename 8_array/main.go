package main

import "fmt"

func main() {
	var fourNum [4]int
	fourNum[0] = 1
	fourNum[1] = 3

	a := [...]int{1, 3, 5}

	fmt.Println("fourNum array: ", fourNum)
	fmt.Println("auto fit size array: ", a)

}
