package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	println("Começou")

	go func() {
		for {
		}
	}()

	time.Sleep(time.Second)
	println("Terminou")
}
