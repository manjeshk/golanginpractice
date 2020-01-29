package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "functions"
	author := "Manjesh"

	//Basic Function
	fmt.Println(converter(module, author))

	//Variadic Function a.k.a var args
	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16)
	fmt.Println(bestFinish)

	sum := addNumbers(1, 2, 3, 4, 5)
	fmt.Println(sum)

}

func bestLeagueFinishes(finishes ...int) int {

	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}

/**
This is an example of basic function
*/
func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)
	return module, author
}

/**
Variadic Function
*/
func addNumbers(numbers ...int) int {
	sum := 0
	for _, i := range numbers {
		sum = sum + i
	}
	return sum
}
