package main

import "time"

func worker(workerId int, msg chan int) {
	for res := range msg {
		println("Worker: ", workerId, " Msg: ", res)
		time.Sleep(time.Microsecond * 250)
	}
}

func main() {
	msg := make(chan int)

	workerSize := 500

	for i := 0; i < workerSize; i++ {
		go worker(i+1, msg)
	}

	ONE_MILLION_WORKS := 1000000
	for i := 0; i < ONE_MILLION_WORKS; i++ {
		msg <- i
	}
}
