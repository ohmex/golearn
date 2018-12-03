package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readAndForward(ch <-chan int64) <-chan int64 {

	r := make(chan int64)

	go func() {
		for val := range ch {
			r <- val
		}
	}()

	return r
}

func main() {
	fmt.Println("my first go routine")
	writech := make(chan int64)
	readch := readAndForward(writech)

	go func() {
		time.Sleep(time.Second * 3)
		writech <- rand.Int63n(100)
	}()

	fmt.Println(<-readch)
}
