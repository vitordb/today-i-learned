package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	arquivo, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer arquivo.Close()

	msg := 2324

	jsonData, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	_, err = arquivo.Write(jsonData)
	if err != nil {
		panic(err)
	}

	fil, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fil))

	fileStats, err := os.Stat("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(fileStats.Size())

	fmt.Println(fileStats.Name())
	fmt.Println(fileStats.IsDir())

}
