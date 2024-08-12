package main

func main() {

	canal := make(chan string, 2)

	go func() {
		canal <- "Ping"
	}()

	go func() {
		canal <- "Pong"
	}()

	println(<-canal)
	println(<-canal)

}
