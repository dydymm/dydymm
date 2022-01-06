package main

import (
	"fmt"
)

func main() {
	person := map[int]string{
		1 : "Tom",
		2 : "Aaron",
		3 : "John",
	}

	fmt.Printf("len=%d map=%v\n", len(person), person)
//其他与数组一样
	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}
}