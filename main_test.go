package main


import (
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