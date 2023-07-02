package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)

	go func() {
		i := 0
		for {
			queue <- i
			i++
		}
	}()

	for value := range queue {
		time.Sleep(time.Second)
		fmt.Printf("foi o i=%d\n", value)
	}
}
