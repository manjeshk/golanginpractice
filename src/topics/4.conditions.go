package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rank1 := "39"
	rank2 := "614"

	//if else condition
	if rank1 < rank2 {
		fmt.Println("First rank")
	} else if rank1 > rank2 {
		fmt.Println("Second rank")
	} else {
		fmt.Println("First rank and Second rank")
	}

	//Simple Initialization Statements
	if rank1, rank2 := 40, 650; rank1 < rank2 {
		fmt.Println("First rank")
	} else if rank1 > rank2 {
		fmt.Println("Second rank")
	} else {
		fmt.Println("First rank and Second rank")
	}

	/*
		Switch case
		1. Implicit breaks
		2. fallthrough execute immediately below cases only
	*/
	switch "docker" {
	case "linux":
		fmt.Println("Linux")
	case "docker":
		fmt.Println("Docker")
		//fallthrough
	case "windows":
		fmt.Println("Windows")
	case "mac":
		fmt.Println("Mac")
	default:
		fmt.Println("Unknown")
	}

	//Second flavor
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Even: ", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd: ", tmpNum)
	}

	/**
	Error Handling
	Idiomatic to return an error as the last return from functions and methods
	nil is used to indicate success
	*/
	//The role of if in Error handling
	//func testFun(target string) (rspTime flowat64 err error)

	_, err := os.Open("/Users/msavita/WorkingDrive/PZN/projects/2019/interstial/test.txt")

	if err != nil {
		fmt.Println("Error occured: ", err)
	}

}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
