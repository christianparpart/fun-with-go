package main

import (
	"fmt"
	"log"

	"github.com/christianparpart/fun-with-go/protobuf-stuff/stuff"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	person := &stuff.Person{
		Id:   proto.Int64(42),
		Name: proto.String("trapni"),
		Email: []string{
			"trapni@gmail.com",
			"trapni@dawanda.com",
		},
	}

	bytes, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Failed to marshal msg. %v\n", err)
	} else {
		fmt.Printf("%+v\n", bytes)
	}

	fmt.Printf("ID: %v\n", person.GetId())
	fmt.Printf("Name: %v\n", person.GetName())
	for i, email := range person.GetEmail() {
		fmt.Printf("Email[%v]: %v\n", i, email)
	}

	m := jsonpb.Marshaler{}
	json, err := m.MarshalToString(person)
	fmt.Printf("json encoded: %v\n", json)

}
