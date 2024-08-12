package main

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {

	var soma T

	for _, v := range m {
		soma += v
	}

	return soma

}

func Compare[T Number](a T, b T) bool {

	return false

}

func main() {

}
