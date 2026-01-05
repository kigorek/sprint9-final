package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"Zero size", 0, 0},
		{"Negative size", -10, 0},
		{"Size 1", 1, 1},
		{"Size 10", 10, 10},
		{"Size 100", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			assert.Equal(t, tt.expected, len(result), "slice length wrong")
			if tt.size > 0 {
				for i, val := range result {
					assert.GreaterOrEqual(t, val, 0, "element at index %d is negative: %d", i, val)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{42}, 42},
		{"Same", []int{5, 5, 5, 5}, 5},
		{"Random", []int{3, 1, 4, 1, 5, 9, 2, 6}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			assert.Equal(t, tt.expected, result)
		})
	}
}
