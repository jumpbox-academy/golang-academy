package main

import "fmt"

func main() {
	// m := make(map[string]string)

	var declareMap map[string]string
	if declareMap == nil {
		fmt.Println("declareMap is nil")
	}

	// m := map[string]string{}

	//declare with initial values
	m := map[string]string{
		"a": "apple",
		"b": "bird",
	}

	m["key"] = "apple"
	m["new-key"] = "jumpbox"

	fmt.Println("map: ", m)
	delete(m, "key")
	fmt.Println("map after delete by key: ", m)

	//Print single value by key
	fmt.Println("m[\"new-key\"]: ", m["new-key"])
}
