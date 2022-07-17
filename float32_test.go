package xboct

import (
	"fmt"
	"testing"
)

func TestFractionalPartLength(t *testing.T) {
	tests := []struct {
		number               float32
		fractionalPartLength uint16
	}{
		//{5},
		//{5.5},
		{5.75, 2},
		{5.74, 19}, // actual number is 5.7399997711181640625
		{5.25, 2},
		{5.5, 1},
		{-5.5, 1},
		{-5.000000476837158203125, 21},
		{5.500000476837158203125, 21},
		{-5.375, 3},
		{-8300000.5, 1},
		{-830000000.5, 0},
		{8400000.5, 0},
		{0, 0},
		{0.5, 1},
		{0.05, 1},
		{0.25, 2},
		{1, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.number), func(t *testing.T) {
			length := FractionalPartLength(tt.number)
			if length != tt.fractionalPartLength {
				t.Errorf("expected length: %d, actial: %d", tt.fractionalPartLength, length)
			}
		})
	}
}
