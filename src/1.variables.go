package main

import (
	"fmt"
	"reflect"
)

//Package level variables are global
var (
	id, name string
	course   string
	module   float64
)

//Infertype
var (
	subject, score = "golang", 90.50
)

func main() {
	fmt.Println("Name is ", name, "and is type", reflect.TypeOf(name))
	fmt.Println("Course is ", course)
	fmt.Println("Module is ", module, "and is of type", reflect.TypeOf(module))

	fmt.Println("------------Infertype-----------------")
	fmt.Println("Module is ", subject, "and is of type", reflect.TypeOf(subject))
	fmt.Println("Module is ", score, "and is of type", reflect.TypeOf(score))

	//--------------------------------Determining types-----------------------------------

	//Shorthand declare and initialize ":=" only works inside of functions
	a := 10.00000
	b := 3

	/*
		The variable is declared but not used. It will throw runtime error because golang doesn't waist the memory.
		The variable declared at function level must be used
	*/
	//d := 10

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of tyoe", reflect.TypeOf(b))

	//c:=a+b
	c := int(a) + b

	fmt.Println("\nC has value:", c, "and is of type:", reflect.TypeOf(c))

}
