package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	concurrency()

}

func concurrency() {
	//enabled parallel
	runtime.GOMAXPROCS(2)

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() { //anonymous function
		defer waitGrp.Done()
		//time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}() //() is used to call self func

	go func() {
		defer waitGrp.Done()
		fmt.Println("Hi")
	}()

	waitGrp.Wait()
}
