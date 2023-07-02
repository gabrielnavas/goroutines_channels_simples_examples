package main

func main() {
	forever := make(chan string)

	go func() {
		for {
		}
	}()

	println("Aguardando para sempre!")

	<-forever
}
