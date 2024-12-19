package main

import (
	"fmt"
)

// values é o array com os números, left é o índice da esquerda, right índice da direita
func partition(values []int, left, right int) int {

	// pega o valor do índice da esquerda como pivot
	pivot := values[left]

	// cria uma variável chamada i que representa a posição onde o valor menor ou igual ao pivot será colocado
	i := left

	// j parte da posição +1 pois a primeira posição é o pivot, enquanto o índice j for menor ou igual que o right, faz a iteração
	for j := left + 1; j <= right; j++ {

		// se o valor for menor ou igual que o pivot
		if values[j] <= pivot {

			// posição é incrementada
			i++
			// troca os valores das posições
			swap(values, i, j)
		}
	}

	// troca o valor do pivot (left) com a última posição que foi incrementada
	swap(values, left, i)

	return i
}

func swap(values []int, i, j int) {
	// troca os valores das posições
	values[i], values[j] = values[j], values[i]
}

// Função recursiva de QuickSort
func quicksort(values []int, left, right int) {
	if left < right {
		// Particiona o array e obtém o índice do pivot
		pivotIndex := partition(values, left, right)

		// Chama recursivamente para o lado esquerdo do pivot
		quicksort(values, left, pivotIndex-1)

		// Chama recursivamente para o lado direito do pivot
		quicksort(values, pivotIndex+1, right)
	}
}

func main() {
	// Array de exemplo
	arr := []int{9, 3, 8, 2, 7, 1, 5}
	arr2 := []int{9, 3, 8, 2, 7, 1, 5}

	// Chama a função quicksort
	quicksort(arr, 0, len(arr)-1)

	// Exibe o array ordenado
	fmt.Println("Array ordenado:", arr)

	fmt.Println("Array ordenado 2:", arr2)

}
