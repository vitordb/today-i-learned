package main

import "fmt"

func main() {

	data := map[string]any{
		"name": "Vitor",
		"age":  31,
	}

	if str, ok := data["name"].(string); ok {

		fmt.Println("value is a string", str)
	}

	if str, ok := data["age"].(string); !ok {
		fmt.Println("not a string", str)
	}
}
