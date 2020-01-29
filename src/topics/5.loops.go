package main

/**
Syntax:
for <expression> {
}
example:
for { //infinite loop
<code>
}

1. blank expression = infinite loop
2. Can be boolean or range
for i <10 {
<code>
}

for i:=range courseList {
<code>
}

for i:= 0; i<10; i++ {
<code>
}

*/

import (
	"fmt"
	"time"
)

func main() {

	for timer := 10; timer >= 0; timer-- {

		if timer == 0 {
			fmt.Println("Boom!")
			break
		}

		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	//range
	courseInProg := []string{"Docker", "Java", "GoLang"}

	for _, i := range courseInProg {
		fmt.Println(i)
	}

}
