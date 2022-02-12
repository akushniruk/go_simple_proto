package main

import (
	"fmt"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{1, 4, 7, 9},
	}

	fmt.Println(sm)
}
