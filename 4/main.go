package main

import "time"

func main() {
	hello := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		hello <- "hello world"
	}()

	select {
	case x := <-hello:
		println(x)
	default:
		println("default")
	}

	time.Sleep(time.Second * 5)
}
