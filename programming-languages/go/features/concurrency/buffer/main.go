package main

import (
	"fmt"
	"time"
)

func main() {
	logChannel := make(chan string, 100) // Canal com buffer

	go func() {
		for logMsg := range logChannel {
			fmt.Println("Logging:", logMsg) // Simulando escrita de log
			time.Sleep(1 * time.Second)     // Simulando delay na escrita
		}
	}()

	// Simulação de múltiplas goroutines enviando logs
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 10; j++ {
				logChannel <- fmt.Sprintf("Goroutine %d - Log %d", id, j)
			}
		}(i)
	}

	time.Sleep(11 * time.Second)
	close(logChannel)
}
