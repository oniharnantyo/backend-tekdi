package unittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Area
func TestSquare_Area(t *testing.T) {
	// Table driven test
	tests := []struct {
		Name          string
		Length        int
		Expect        int
		IsExpectError bool
	}{
		{
			Name:          "square area error",
			Length:        10,
			Expect:        50,
			IsExpectError: true,
		},
		{
			Name:          "square area no error",
			Length:        10,
			Expect:        100,
			IsExpectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			square := NewSquare(10)

			expectResult := test.Expect

			length := square.Area()

			if test.IsExpectError {
				assert.NotEqual(t, expectResult, length)
			} else {
				assert.Equal(t, expectResult, length)
			}
		})
	}
}

func TestSquare_AreaError(t *testing.T) {
	square := NewSquare(10)

	expectResult := 50

	length := square.Area()

	assert.NotEqual(t, expectResult, length)
}

func TestSquare_AreaNoError(t *testing.T) {
	square := NewSquare(10)

	expectResult := 100

	length := square.Area()

	assert.Equal(t, expectResult, length)
}

// Around
func TestSquare_AroundError(t *testing.T) {
	square := NewSquare(10)

	expectResult := 50

	around := square.Around()

	assert.NotEqual(t, expectResult, around)
}

func TestSquare_AroundNoError(t *testing.T) {
	square := NewSquare(10)

	expectResult := 40

	around := square.Around()

	assert.Equal(t, expectResult, around)
}
