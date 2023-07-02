package main

// Thread 1
func main() {

	// Thread 1 <-> Thread 2
	hello := make(chan string)

	// Thread 2
	go func() {
		hello <- "Hello World"
	}()

	result := <-hello
	println(result)

}
