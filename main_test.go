package main


import (
	"math/rand"
	"testing"

    //"github.com/stretchr/testify/require"
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
		maxValue := rand.Intn(MAX)
		slice[index] = maxValue
		assert.Equal(t, maximum(slice), maxValue)
	}
	assert.Equal(t, maximum([]int{}), 0)
}