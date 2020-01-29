package main

import (
	"fmt"
)

func main() {

	//Create map
	titles := make(map[string]int)
	titles["title1"] = 6
	titles["title2"] = 4

	anotherWay := map[string]int{
		"title1": 5,
		"title2": 3,
	}

	fmt.Printf("Title1: %v and anotherWay: %v", titles, anotherWay)

	//Iterate map
	testMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	//It will print the entry in random order because it uses random offset
	for key, value := range testMap {
		fmt.Printf("\nKey: %v and value: %v", key, value)
	}

	fmt.Println("\n", testMap["a"])
	testMap["a"] = 100
	fmt.Println(testMap["a"])

	delete(testMap, "e")
	fmt.Println(testMap)

	//maps are passed by reference rather than value

}
