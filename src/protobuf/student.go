package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {

	student := &Student{
		Name: "John",
		Age:  24,
		Address: &Address{
			Street: "1410 Pioneer St",
			City:   "Phoenix",
			State:  "Arizona",
			Zip:    85027,
		},
	}

	data, err := proto.Marshal(student)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	fmt.Println(data)

	newStudent := &Student{}
	err = proto.Unmarshal(data, newStudent)
	if err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}

	fmt.Println(newStudent.Name)
	fmt.Println(newStudent.Age)
	fmt.Println(newStudent.Address.Street)
	fmt.Println(newStudent.Address.City)
	fmt.Println(newStudent.Address.State)
	fmt.Println(newStudent.Address.Zip)

}
