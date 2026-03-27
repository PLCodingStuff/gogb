package main

import "testing"

func TestNOP(t *testing.T) {

}

func TestJP_a16(t *testing.T) {

}

func TestJP_a16_LittleEndian(t *testing.T) {

}

func TestLD_A_d8(t *testing.T) {

}

func TestLDH_a8_A(t *testing.T) {

}

func TestHALT(t *testing.T) {

}

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
