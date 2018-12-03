package main

import (
	"fmt"
)

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

func newIntByKeyword() *int {
	return new(int)
}

// Both are same
func newIntByVar() *int {
	var dummy int
	return &dummy
}
