package main

const N = 3

func main() {
	m := make(map[int]*int)

	for i := 0; i < N; i++ {
		m[i] = &i
	}

	for _, v := range m {
		println(*v)
	}
}
