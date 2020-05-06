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

// TestRemoveDuplicateMousePoints tests the RemoveDuplicateMousePoints function
func TestRemoveDuplicateMousePoints(t *testing.T) {
	cases := []struct {
		given    []NormalizedMouseData
		expected []NormalizedMouseData
		name     string
	}{
		{
			given:    []NormalizedMouseData{{X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			expected: []NormalizedMouseData{{X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			name:     "Two Double Points at beginning",
		},
		{
			given:    []NormalizedMouseData{{X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			expected: []NormalizedMouseData{{X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			name:     "Three Double Points at beginning",
		},
		{
			given:    []NormalizedMouseData{{X: 1, Y: 1, Time: 5}},
			expected: []NormalizedMouseData{{X: 1, Y: 1, Time: 5}},
			name:     "Only one point",
		},
		{
			given:    []NormalizedMouseData{{X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			expected: []NormalizedMouseData{{X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			name:     "Two points at the same time, not next to each other",
		},
		{
			given:    []NormalizedMouseData{{X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			expected: []NormalizedMouseData{{X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			name:     "Two Double Points in the middle",
		},

		{
			given:    []NormalizedMouseData{{X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			expected: []NormalizedMouseData{{X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 6}},
			name:     "Two Double Points at the beginning and in the middle",
		},
		{
			given:    []NormalizedMouseData{{X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 5}},
			expected: []NormalizedMouseData{{X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 5}},
			name:     "Two Double Points at the end",
		},
		{
			given:    []NormalizedMouseData{{X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 5}},
			expected: []NormalizedMouseData{{X: 1, Y: 1, Time: 4}, {X: 1, Y: 1, Time: 5}},
			name:     "Two Double Points at the beginning and end",
		},
	}

	for _, test := range cases {
		if res := RemoveDuplicateMousePoints(test.given); !equalPoints(res, test.expected) {
			t.Errorf("%v failed: %v given, expected %v\n", test.name, res, test.expected)
		}
	}
}

// equalPoints checks if two arrays are the same
func equalPoints(points, otherPoints []NormalizedMouseData) bool {
	if len(points) != len(otherPoints) {
		return false
	}
	for i := range points {
		if !points[i].Equals(otherPoints[i]) {
			return false
		}
	}
	return true
}

// equalScrolls checks if two arrays are the same
func equalScrolls(scrolls, otherScrolls []NormalizedScrollEvent) bool {
	if len(scrolls) != len(otherScrolls) {
		return false
	}
	for i := range scrolls {
		if !scrolls[i].Equals(otherScrolls[i]) {
			return false
		}
	}
	return true
}

// TestRemoveConcurrentEvents tests the RemoveConcurrentEvent function
func TestRemoveConcurrentEvents(t *testing.T) {
	cases := []struct {
		name            string
		clicks          []NormalizedClickEvent
		points          []NormalizedMouseData
		scrolls         []NormalizedScrollEvent
		expectedPoints  []NormalizedMouseData
		expectedScrolls []NormalizedScrollEvent
	}{
		{
			name:            "No concurrent Events",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 10}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
		},
		{
			name:            "Concurrent Move and Scroll removes Scroll",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 10}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 6}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
		}, {
			name:            "Concurrent Click and Move removes Move",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 10}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 8}},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 8}},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
		}, {
			name:            "Concurrent Click and Scroll removes Scroll",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 10}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
		},
		{
			name:            "Concurrent Click and Move and Scroll removes Move and Scroll",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 10}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 5}, {X: 1, Y: 1, Time: 8}},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 8}},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
		},
		{
			name:            "Unordered Concurrency with Click and the others",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 4}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 4}},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 5}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 6}},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}},
		},
		{
			name:            "Unordered Concurrency between Point and Scroll",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 35}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 10}},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 10}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 6}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 11}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 10}},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 11}, 1, 2, 3, 0}},
		},
		{
			name:            "Empty Scrolls no error",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 10}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			scrolls:         []NormalizedScrollEvent{},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			expectedScrolls: []NormalizedScrollEvent{},
		},
		{
			name:            "Empty Points no error",
			clicks:          []NormalizedClickEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 5}, MouseKeyLEFT, false}, {NormalizedMouseData{X: 1, Y: 1, Time: 10}, MouseKeyLEFT, false}},
			points:          []NormalizedMouseData{},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
		},
		{
			name:            "Empty Clicks no error",
			clicks:          []NormalizedClickEvent{},
			points:          []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			scrolls:         []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
			expectedPoints:  []NormalizedMouseData{{X: 1, Y: 1, Time: 6}, {X: 1, Y: 1, Time: 8}},
			expectedScrolls: []NormalizedScrollEvent{{NormalizedMouseData{X: 1, Y: 1, Time: 7}, 1, 2, 3, 0}, {NormalizedMouseData{X: 1, Y: 1, Time: 9}, 1, 2, 3, 0}},
		},
	}

	for _, test := range cases {
		points, scrolls := RemoveConcurrentEvents(test.clicks, test.points, test.scrolls)
		if !equalPoints(points, test.expectedPoints) || !equalScrolls(scrolls, test.expectedScrolls) {
			t.Errorf("%v failed: Excpected %v and %v, but got %v and %v\n", test.name, test.expectedPoints, test.expectedScrolls, points, scrolls)
		}
	}
}
