package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rad() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- int(rand.New(rand.NewSource(time.Now().UnixNano())))
		}
	}()
	return ch
}
func main() {
	ch := rad()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
