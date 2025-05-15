package main

import "fmt"

func main() {
	aMap := make(map[string]struct{})

	aMap["a"] = struct{}{}

	for key, _ := range aMap {
		fmt.Print(key)
	}
}
