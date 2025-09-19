package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
	MAXVAL = 10_000
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <=0 {
		return []int{}
	}
	result := make([]int, size)
	for i:=0; i<size; i++ {
		item := rand.Intn(MAXVAL)
		result[i] = item
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	result := data[0]
	for i:=1; i<len(data); i++ {
		if result < data[i] {
			result = data[i]
		}
	}
	return result
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	if len(data)<8{
		return maximum(data)
	}
	ch := make(chan int)
	chunkLen := len(data) / CHUNKS
	for i:=0; i<CHUNKS; i++ {
		startIndex := i * chunkLen
		endIndex := startIndex + chunkLen
		if i == CHUNKS - 1 {
			endIndex = len(data)
		}
		wg.Add(1)
		go func(data []int, startIndex int, endIndex int, ch chan int){
			result := maximum(data[startIndex:endIndex])
			ch <- result
			wg.Done()
		}(data, startIndex, endIndex, ch)
	}

	maxSlice := make([]int, 0)
	for i :=0; i<8; i++ {
		maxSlice = append(maxSlice, <-ch)
	}
	wg.Wait()
	return maximum(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Now().Sub(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	start = time.Now()
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	max = maxChunks(numbers)
	elapsed = time.Now().Sub(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
