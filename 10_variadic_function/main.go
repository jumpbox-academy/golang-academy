package main

import "fmt"

func main() {
	variadicSample("Tech Passion", "Sharing", "Society")

	// or
	// s := []string{"Tech Passion", "Sharing", "Society"}
	// variadicSample(s...)
}

func variadicSample(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println("Jumpbox Pillar ", i+1, ": ", s[i])
	}
}
