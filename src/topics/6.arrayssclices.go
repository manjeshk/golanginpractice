package main

import (
	"fmt"
)

func main() {

	//Declare a slice called myCources
	myCources := make([]string, 5, 10)

	fmt.Printf("Length is: %d. \nCapacity is %d", len(myCources), cap(myCources))

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice[4])

	sliceOfSlice1 := mySlice[2:5]
	fmt.Println(sliceOfSlice1)

	sliceOfSlice2 := mySlice[:5]
	fmt.Println(sliceOfSlice2)

	sliceOfSlice3 := mySlice[2:]
	fmt.Println(sliceOfSlice3)

	//append or expand the slice
	mySlice = append(mySlice, 11)
	mySlice = append(mySlice, 12)
	fmt.Println(mySlice)

	newSlice := []int{12, 13, 14, 15}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

}
