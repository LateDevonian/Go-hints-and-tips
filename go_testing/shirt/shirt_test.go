package tshirtcalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePrice(t *testing.T) {
	// We define a struct inlined as it won't be used outside of this test function
	testCases := []struct {
		name string

		// inputs
		color Color
		size  Size

		// expected outputs
		expectedPrice float64
	}{
		// Declaring a struct the long/normal way
		{
			name:          "Beige S",
			size:          S,
			expectedPrice: 3.50,
		},
		// This is a short-hand where you don't need to specift struct fields.
		{"Beige S", Beige, S, 3.50},
		{"Beige M", Beige, M, 3.50},
		{"Beige L", Beige, L, 3.50},

		{"Khaki S", Khaki, S, 3.50},
		{"Khaki M", Khaki, M, 3.50},
		{"Khaki L", Khaki, L, 3.50},

		{"Cream S", Cream, S, 3.50},
		{"Cream M", Cream, M, 3.50},
		{"Cream L", Cream, L, 3.50},

		{"CosmicLatte S", CosmicLatte, S, 3.50},
		{"CosmicLatte M", CosmicLatte, M, 3.50},
		{"CosmicLatte L", CosmicLatte, L, 3.50},

		{"UnbleachedSilk S", UnbleachedSilk, S, 3.50},
		{"UnbleachedSilk M", UnbleachedSilk, M, 3.50},
		{"UnbleachedSilk L", UnbleachedSilk, L, 3.50},

		// Task 6: Add test cases for Blue T-shirts
		{"Blue S", Blue, S, 13.50},
		{"Blue S", Blue, M, 13.50},
		{"Blue S", Blue, L, 13.50},

		// Task 7: Add/update test cases for the text printing fee
	}

	for _, tc := range testCases {
		// Each test case is effectively a sub-test
		t.Run(tc.name, func(t *testing.T) {
			actualPrice := CalculatePrice(tc.size, tc.color, "")

			assert.Equal(t, tc.expectedPrice, actualPrice)
		})
	}
}

// Not working? Run it in browser: https://goplay.space/#9UYDVxiqxVX
