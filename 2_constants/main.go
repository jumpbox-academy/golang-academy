package main

import "fmt"

func main() {

	const INT_VALUE = 6
	const STRING_VALUE = "Jumpbox"
	const BOOL_VALUE = true

	const (
		sunday = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	fmt.Println("====  constants ====")
	fmt.Println("INT_VALUE: ", INT_VALUE)
	fmt.Println("STRING_VALUE: ", STRING_VALUE)
	fmt.Println("BOOL_VALUE: ", BOOL_VALUE, "\n")

	fmt.Println("==== Enum with constant ====")
	fmt.Println("enum - monday: ", monday)
	fmt.Println("enum - tuesday: ", tuesday)
	fmt.Println("enum - wednesday: ", wednesday)

}
