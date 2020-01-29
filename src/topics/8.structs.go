package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	DockerDeepDive := courseMeta{
		Author: "Nigel",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println(DockerDeepDive.Author)

}
