package main

import "fmt"

type book struct {
	name   string
	author string
}

func main() {
	var cloudNative book = book{
		name:   "Cloud Native Observability with OpenTelemetry",
		author: "Alex Boten",
	}
	fmt.Println("book name: ", cloudNative.name)
	fmt.Println("author: ", cloudNative.author)
}
