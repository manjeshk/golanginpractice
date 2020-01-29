package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("-------------------Pointers-------------")
	course := "Go Lang"
	module := 3.2

	//Create pointer variable
	ptr := &module

	fmt.Println("Course is", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Memory address of *module variable is", &module)

	/*
		& - References a pointer
		* - De-references a pointer
	*/

	fmt.Println("Memory address of *module* variable is", &ptr, "and the value of *module* is", *ptr)

	fmt.Println("-------------------Passing By Values-------------")
	fmt.Println("The course is ", course)
	changeCourseByValue(course)
	fmt.Println("The course is ", course)

	fmt.Println("-------------------Passing By Reference-------------")
	fmt.Println("The course is ", course)
	changeCourseByReference(&course)
	fmt.Println("The course is ", course)

	fmt.Println("-------------------Constant-------------")
	const name = "Manjesh"

}

func changeCourseByValue(course string) string {
	course = "Machine Learning"
	fmt.Println("The changed course is ", course)

	return course
}

func changeCourseByReference(course *string) string {
	*course = "Machine Learning"
	fmt.Println("The changed course is ", *course)

	return *course
}
