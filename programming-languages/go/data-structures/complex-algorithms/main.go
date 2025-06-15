package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Representa a distância entre cidades
var distances = map[string]map[string]int{
	"A": {"B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6},
	"B": {"C": 1, "D": 2, "E": 3, "F": 4, "G": 5},
	"C": {"D": 1, "E": 2, "F": 3, "G": 4},
	"D": {"E": 1, "F": 2, "G": 3},
	"E": {"F": 1, "G": 2},
	"F": {"G": 1},
}

// Calcula a distância total para uma sequência de cidades
func calculateDistance(route []string) int {
	total := 0
	for i := 0; i < len(route)-1; i++ {
		from, to := route[i], route[i+1]
		if d, ok := distances[from][to]; ok {
			total += d
		} else if d, ok := distances[to][from]; ok {
			total += d
		} else {
			total += 1000 // Penalidade para rotas inválidas
		}
	}
	return total
}

// Realiza mutação trocando duas cidades de posição
func mutate(route []string) []string {
	mutated := make([]string, len(route))
	copy(mutated, route)
	i, j := rand.Intn(len(route)), rand.Intn(len(route))
	mutated[i], mutated[j] = mutated[j], mutated[i]
	return mutated
}

// Realiza crossover entre dois slices
func crossover(parent1, parent2 []string) []string {
	child := make([]string, 0, len(parent1))
	used := make(map[string]bool)

	// Pegar a primeira metade do primeiro pai
	for i := 0; i < len(parent1)/2; i++ {
		child = append(child, parent1[i])
		used[parent1[i]] = true
	}

	// Completar com o segundo pai
	for _, city := range parent2 {
		if !used[city] {
			child = append(child, city)
		}
	}

	return child
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Cidades iniciais
	cities := []string{"A", "B", "C", "D", "E", "F", "G"}
	bestRoute := make([]string, len(cities))
	copy(bestRoute, cities)

	// Embaralhar as cidades para gerar uma sequência inicial
	rand.Shuffle(len(bestRoute), func(i, j int) { bestRoute[i], bestRoute[j] = bestRoute[j], bestRoute[i] })

	bestDistance := calculateDistance(bestRoute)
	fmt.Printf("Initial Route: %v, Distance: %d\n", bestRoute, bestDistance)

	// Loop infinito até atingir um limite de tentativas ou encontrar o menor caminho
	for iteration := 1; ; iteration++ {
		// Mutação
		mutatedRoute := mutate(bestRoute)
		mutatedDistance := calculateDistance(mutatedRoute)

		// Crossover
		crossedRoute := crossover(bestRoute, mutatedRoute)
		crossedDistance := calculateDistance(crossedRoute)

		// Log de cada tentativa
		fmt.Printf("Iteration %d:\n", iteration)
		fmt.Printf("  Mutated Route: %v, Distance: %d\n", mutatedRoute, mutatedDistance)
		fmt.Printf("  Crossed Route: %v, Distance: %d\n", crossedRoute, crossedDistance)

		// Seleção: verificar se a distância melhorou
		if mutatedDistance < bestDistance {
			bestRoute = mutatedRoute
			bestDistance = mutatedDistance
		}
		if crossedDistance < bestDistance {
			bestRoute = crossedRoute
			bestDistance = crossedDistance
		}

		// Condição de parada (exemplo: número máximo de iterações ou distância mínima)
		if iteration >= 10000 || bestDistance <= 10 {
			fmt.Printf("Iteration %d: Best Route Found: %v, Distance: %d\n", iteration, bestRoute, bestDistance)
			break
		}
	}
}
