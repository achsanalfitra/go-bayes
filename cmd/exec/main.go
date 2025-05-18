package main

import "fmt"

func main() {
	aMap := make(map[string]map[string]struct{})
	// if _, ok := aMap["water"]; !ok {
	// 	aMap["water"] = make(map[string]struct{})
	// }

	// aMap["water"]["cooled"] = struct{}{}

	// aMap["chill"] = map[string]struct{}{}

	// if _, exists := aMap["water"]; !exists {
	// 	fmt.Println("water doesn't exist")
	// } else {
	// 	fmt.Println("water does exist")
	// }

	if _, exists := aMap["chill"]; !exists {
		fmt.Println("chill doesn't exist")
	} else {
		fmt.Println("chill does exist")
	}

	if _, exists := aMap["water"]["cooled"]; !exists {
		fmt.Println("water cooled doesn't exist")
	} else {
		fmt.Println("water cooled does exist")
	}

	_, assignCheck := aMap["water"]["cooled"]
	if assignCheck {
		fmt.Println("assignCheck works")
	} else {
		fmt.Println("assignCheck works")
	}

	fmt.Println(len(aMap))
}
