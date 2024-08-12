package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("large_file.txt")
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte("Hello, World!\nThis is a test file"))
	if err != nil {
		panic(err)
	}

	defer file.Close()

	bufferSize := 1024
	buffer := make([]byte, bufferSize)

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF { // Fim do arquivo
				break
			}
			fmt.Println("Erro ao ler o arquivo:", err)
			return
		}

		// Processa os bytes lidos aqui
		fmt.Println("Bytes lidos:", bytesRead)
		fmt.Println(string(buffer[:bytesRead])) // Convertendo os bytes lidos em string para demonstração
	}
}
