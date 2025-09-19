package main


import (
	"math/rand"
	"testing"

    "github.com/stretchr/testify/assert"
)


func TestGenerateRandomElements(t *testing.T) {
	for size := 0; size < 1000; size++ {
		randSlice := generateRandomElements(size)
		assert.Len(t, randSlice, size)
	}
}

func TestMaximum(t *testing.T){
	for size:=1; size<1000; size++{
		slice := make([]int, size)
		index := rand.Intn(size)
		maxValue := rand.Intn(MAXVAL)
		slice[index] = maxValue
		assert.Equal(t, maximum(slice), maxValue)
	}
	assert.Equal(t, maximum([]int{}), 0)
}

func TestMaxChunks(t *testing.T) {
	for _, size:=range []int{7, 8, 9, 70, 80, 90, 1000, 1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 2000, 3000}{
		slice := make([]int, size)
		index := rand.Intn(size)
		maxValue := rand.Intn(MAXVAL)
		slice[index] = maxValue
		assert.Equal(t, maxValue, maxChunks(slice))
	}
	assert.Equal(t, maximum([]int{}), 0)
}