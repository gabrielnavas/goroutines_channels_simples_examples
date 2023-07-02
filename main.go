package main

import "time"

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go contador("a")
	go contador("b")
	time.Sleep(time.Second * 10)
}
