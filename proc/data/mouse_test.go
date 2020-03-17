package data

import "testing"

// TestNormalizeMouse tests the mouse.Normalize function
func TestNormalizeMouse(t *testing.T) {
	var cases = []struct {
		name      string
		point     MouseData
		screen    ScreenInfo
		startTime uint64
		expected  NormalizedMouseData
	}{
		{"full size", MouseData{X: 100, Y: 70, Time: 150}, ScreenInfo{Width: 100, Height: 70}, 10, NormalizedMouseData{X: 1, Y: 1, Time: 140}},
		{"null", MouseData{X: 0, Y: 0, Time: 10}, ScreenInfo{Width: 100, Height: 70}, 10, NormalizedMouseData{X: 0, Y: 0, Time: 0}},
		{"middle", MouseData{X: 50, Y: 35, Time: 160}, ScreenInfo{Width: 100, Height: 70}, 10, NormalizedMouseData{X: 0.5, Y: 0.5, Time: 150}},
	}

	for _, c := range cases {
		output := c.point.Normalize(c.screen, Time{Start: c.startTime})
		if output != c.expected {
			t.Errorf("%v failed: %v was expected, but got %v\n", c.name, c.expected, output)
		}
	}
}

// TestNormalizeMouseData tests the NormalizeMouseData function
func TestNormalizeMouseData(t *testing.T) {
	cases := []MouseData{{X: 100, Y: 70, Time: 150}, {X: 0, Y: 0, Time: 10}, {X: 50, Y: 35, Time: 160}}
	expected := []NormalizedMouseData{{X: 1, Y: 1, Time: 140}, {X: 0, Y: 0, Time: 0}, {X: 0.5, Y: 0.5, Time: 150}}
	output := NormalizeMouseData(cases, ScreenInfo{Width: 100, Height: 70}, Time{Start: 10})
	for i, o := range output {
		if o != expected[i] {
			t.Errorf("%v was not %v\n", output, expected)
			break
		}
	}
}
