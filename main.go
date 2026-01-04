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
)

// generateRandomElements generates random positive integers.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	data := make([]int, size)

	for i := 0; i < size; i++ {
		data[i] = rng.Int()
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return maximum(data)
	}

	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		go func(chunkIndex int) {
			defer wg.Done()
			start := chunkIndex * chunkSize
			end := start + chunkSize

			if chunkIndex == CHUNKS-1 {
				end = len(data)
			}
			chunk := data[start:end]

			if len(chunk) > 0 {
				chunkMax := chunk[0]
				for j := 1; j < len(chunk); j++ {
					if chunk[j] > chunkMax {
						chunkMax = chunk[j]
					}
				}
				maxValues[chunkIndex] = chunkMax
			} else {
				maxValues[chunkIndex] = 0
			}
		}(i)
	}
	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)
}
