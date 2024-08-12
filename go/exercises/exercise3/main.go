package main

func main() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	if v, ok := m["b"]; ok {
		println(v)
	}
}
