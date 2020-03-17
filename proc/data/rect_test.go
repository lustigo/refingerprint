package data

import "testing"

// TestNormalizeRect tests the rect.Normalize function
func TestNormalizeRect(t *testing.T) {
	var cases = []struct {
		name     string
		screen   ScreenInfo
		rect     Rectangle
		expected Rectangle
	}{
		{"null", ScreenInfo{Width: 100, Height: 50}, Rectangle{}, Rectangle{}},
		{"without delta", ScreenInfo{Width: 100, Height: 50}, Rectangle{X: 10, Y: 5, Width: 20, Height: 35}, Rectangle{X: 0.1, Y: 0.1, Width: 0.2, Height: 0.7}},
		{"with delta", ScreenInfo{Width: 100, Height: 50, DeltaX: 3, DeltaY: 4}, Rectangle{X: 7, Y: 1, Width: 20, Height: 35}, Rectangle{X: 0.1, Y: 0.1, Width: 0.2, Height: 0.7}},
	}

	for _, c := range cases {
		if c.rect.Normalize(c.screen) != c.expected {
			t.Errorf("%v failed: %v is not %v\n", c.name, c.rect.Normalize(c.screen), c.expected)
		}
	}

}
