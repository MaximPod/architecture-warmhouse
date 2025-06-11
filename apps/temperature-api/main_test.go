package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomTemperature(t *testing.T) {
	tests := []struct {
		name    string
		lowest  float64
		highest float64
		want    bool
	}{
		{"Test with rand.Float64() = 0", -30, 40, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.New(rand.NewSource(rand.Int63()))
			result := getRandomTemperature(rand.New(rand.NewSource(0)))
			got := isInRange(result, tt.lowest, tt.highest)
			if got != tt.want {
				t.Errorf("isInRange(%f, %f, %f) = %v, want %v", result, tt.lowest, tt.highest, got, tt.want)
			}
		})
	}
}

func isInRange(x, min, max float64) bool {
	return x >= min && x <= max
}

func TestRoundFloat(t *testing.T) {
	tests := []struct {
		name      string
		val       float64
		precision uint
		expected  float64
	}{
		{"Round to 0 decimal places", 1.2345, 0, 1},
		{"Round to 1 decimal place", 1.2345, 1, 1.2},
		{"Round to 2 decimal places", 1.2345, 2, 1.23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := roundFloat(tt.val, tt.precision)
			assert.Equal(t, tt.expected, result)
		})
	}
}
