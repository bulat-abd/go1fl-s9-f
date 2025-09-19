package main

import (
	"fmt"
	"math/rand"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
	MAX    = 10_000
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <=0 {
		return []int{}
	}
	result := make([]int, size)
	for i:=0; i<size; i++ {
		item := rand.Intn(MAX)
		result[i] = item
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	return 0
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	return 0
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	max := 0
	elapsed := 0
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
