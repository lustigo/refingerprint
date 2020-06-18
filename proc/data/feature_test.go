package data

import "testing"

// TestConvert1DTo2D tests the TestConvert1DTo2D method
func TestConvert1DTo2D(t *testing.T) {
	var cases = []struct {
		name      string
		posno     uint8
		size      uint8
		expectedX uint8
		expectedY uint8
	}{
		{"Size = 3, Position (0,0)", 0, 3, 0, 0},
		{"Size = 3, Position (1,0)", 1, 3, 1, 0},
		{"Size = 3, Position (2,0)", 2, 3, 2, 0},
		{"Size = 3, Position (0,2)", 6, 3, 0, 2},
		{"Size = 3, Position (2,2)", 8, 3, 2, 2},
		{"Size = 4, Position (0,0)", 0, 4, 0, 0},
		{"Size = 4, Position (1,0)", 1, 4, 1, 0},
		{"Size = 4, Position (3,0)", 3, 4, 3, 0},
		{"Size = 4, Position (0,3)", 12, 4, 0, 3},
		{"Size = 4, Position (3,3)", 15, 4, 3, 3},
	}

	for _, c := range cases {
		x, y := Convert1DTo2D(c.posno, c.size)
		if x != c.expectedX || y != c.expectedY {
			t.Errorf("%v failed: Expected (%v,%v) but got (%v,%v)\n", c.name, c.expectedX, c.expectedY, x, y)
		}
	}
}
