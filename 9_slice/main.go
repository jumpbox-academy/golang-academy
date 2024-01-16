package main

import "fmt"

func main() {
	var number []int
	fmt.Println("zero value number: ", number, "init length: ", len(number), "initial capacity: ", cap(number))

	number = make([]int, 3, 6)

	fmt.Println("after make: ", number, "length: ", len(number), "capacity: ", cap(number))

	// try to print over initial
	// fmt.Println("after make: ", number[3])

	number[0] = 1
	number[1] = 2
	fmt.Println("berfore append: ", number)

	// try to assign over length
	// number[3] = 11
	// fmt.Println("over length: ", number)

	number = append(number, 6)
	fmt.Println("after append: ", number, "length: ", len(number), "capacity: ", cap(number))

	// Scenarios 2 append over capacity
	// number = append(number, 5) // index 5
	// number = append(number, 4) // index 6
	// number = append(number, 3) // over capacity
	// fmt.Println("over capacity: ", number, "length: ", len(number), "capacity: ", cap(number))
	// or
	// number = append(number, 5, 4, 3)

	// Scenarios 3: Array and Slice
	// var array_num [6]int // init Array lenght 6 capacity 6
	// fmt.Println("init array", array_num, "length: ", len(array_num), "capacity: ", cap(array_num))

	// array_num[0] = 3
	// array_num[1] = 2
	// array_num[2] = 4
	// array_num[3] = 5
	// fmt.Println("array_num", array_num, "length: ", len(array_num), "capacity: ", cap(array_num))

	// var slice_num []int        //init Slice
	// slice_num = array_num[0:5] // lenght 2 capacity: go initial
	// fmt.Println("part_number", slice_num, "length: ", len(slice_num), "capacity: ", cap(slice_num))

	// // Cast type Array to Slice
	// slice_num = array_num[0:6]
	// // or
	// slice_num = array_num[:]
}
