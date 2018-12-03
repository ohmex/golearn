package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n = flag.Bool("n", false, "omit trainling newline")
	var sep = flag.String("s", " ", "seperator")

	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

//go run gopl.io/flags/main.go -n -s / a b cd ef
