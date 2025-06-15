package main

func main() {
	m["a"] = 1
	m["b"] = 2

	if v, ok := m["b"]; ok {
		println(v)
	}
}
