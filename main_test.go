package main

import (
	"testing"
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

			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) = slice with length %d, want %d",
					tt.size, len(result), tt.expected)
			}
			if tt.size > 0 {
				for i, val := range result {
					if val < 0 {
						t.Errorf("Element at index %d is negative: %d", i, val)
					}
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
		{"Negative", []int{-5, -1, -3, -10}, -1},
		{"Mix", []int{-5, 0, 10, -3, 7}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, result, tt.expected)
			}
		})
	}
}
