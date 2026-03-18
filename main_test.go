package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"two positives", 2, 3, 5},
		{"two negatives", -2, -3, -5},
		{"positive and negative", 10, -4, 6},
		{"zeros", 0, 0, 0},
		{"one zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
