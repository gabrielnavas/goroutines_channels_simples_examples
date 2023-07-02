package main

import "time"

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		println(tipo, i)
	}
}

func main() {
	contador("sem go routine")
	go contador("com go routine")
	println("Hello 1")
	println("Hello 2")
	time.Sleep(time.Second)
}
