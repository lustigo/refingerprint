package data

import (
	"math"
	"testing"
)

// TestNormalizeRect tests the rect.Normalize function
func TestNormalizeRect(t *testing.T) {
	var cases = []struct {
		name     string
		screen   ScreenInfo
		rect     Rectangle
		expected Rectangle
	}{
		{"null", ScreenInfo{Width: 100, Height: 50}, Rectangle{}, Rectangle{}},
		{"without delta", ScreenInfo{Width: 100, Height: 50}, Rectangle{X: 10, Y: 5, Width: 20, Height: 35}, Rectangle{X: 0.0894, Y: 0.0447, Width: 0.1788, Height: 0.3130}},
		{"with delta", ScreenInfo{Width: 100, Height: 50, DeltaX: 3, DeltaY: 4}, Rectangle{X: 7, Y: 1, Width: 20, Height: 35}, Rectangle{X: 0.0894, Y: 0.0447, Width: 0.1788, Height: 0.3130}},
	}

	for _, c := range cases {
		output := c.rect.Normalize(c.screen)
		if math.Abs(output.X-c.expected.X) > 0.01 || math.Abs(output.Y-c.expected.Y) > 0.01 || math.Abs(output.Width-c.expected.Width) > 0.01 || math.Abs(output.Height-c.expected.Height) > 0.01 {
			t.Errorf("%v failed: %v is not %v\n", c.name, c.rect.Normalize(c.screen), c.expected)
		}
	}

}
