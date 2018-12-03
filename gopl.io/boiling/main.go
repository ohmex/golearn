package main

import (
	"fmt"
)

const temp = 212.0

func main() {
	var f = temp
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %gF or %%%.2fC\n", f, c)
	fmt.Printf("Boiling point = %gF or %.2fC\n", f, f2c(f))
}

func f2c(f float64) float64 {
	return (f - 32) * 5 / 9
}
