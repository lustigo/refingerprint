package data

import (
	"testing"
)

// TestBinSearch tests the BinSearch function
func TestBinSearch(t *testing.T) {
	var cases = []struct {
		name     string
		haystack []string
		needle   string
		expected bool
	}{
		{"no haystack", []string{}, "anyString", false},
		{"no needle", []string{"a", "b", "c"}, "", false},
		{"no needle no haystack", []string{}, "", false},
		{"no needle haystack with empty string", []string{""}, "", true},

		{"needle not in haystack with size 3, would be at the end", []string{"a", "b", "c"}, "d", false},
		{"needle not in haystack with size 3, would be at the front", []string{"b", "c", "d"}, "a", false},
		{"needle not in haystack with size 3, would be in the middle-front", []string{"a", "c", "d"}, "b", false},
		{"needle not in haystack with size 3, would be in the middle-back", []string{"a", "b", "d"}, "c", false},
		{"needle not in haystack with size 4, would be at the end", []string{"a", "b", "c", "d"}, "e", false},
		{"needle not in haystack with size 4, would be at the front", []string{"b", "c", "d", "e"}, "a", false},
		{"needle not in haystack with size 4, would be in the middle", []string{"a", "b", "d", "e"}, "c", false},

		{"needle  in haystack with size 3, is in the end", []string{"a", "b", "c"}, "c", true},
		{"needle in haystack with size 3, is in the front", []string{"b", "c", "d"}, "b", true},
		{"needle in haystack with size 3, is in the middle", []string{"a", "c", "d"}, "c", true},
		{"needle in haystack with size 4, is at the end", []string{"a", "b", "c", "d"}, "d", true},
		{"needle in haystack with size 4, is at the front", []string{"b", "c", "d", "e"}, "b", true},
		{"needle in haystack with size 4, is in the middle-end", []string{"a", "b", "d", "e"}, "d", true},
		{"needle in haystack with size 4, is in the middle-front", []string{"a", "b", "d", "e"}, "b", true},
	}
	for _, c := range cases {
		if BinSearch(c.haystack, c.needle) != c.expected {
			t.Errorf("%v failed\n", c.name)
		}
	}
}

// Generates an ordered string array with the given length
func generateArray(number int) []string {
	s := "a"
	arr := []string{}
	for i := 0; i < number; i++ {
		arr = append(arr, s)
		s = s + "a"
	}
	return arr
}

// BenchmarkLinSearch1 benchmarks a binary search where a needle is not found with 1 element
func BenchmarkBinSearch1(b *testing.B) {
	BinSearch(generateArray(1), "no language")
}

// BenchmarkLinSearch10 benchmarks a binary search where a needle is not found with 10 element
func BenchmarkBinSearch10(b *testing.B) {
	BinSearch(generateArray(10), "no language")
}

// BenchmarkLinSearch100 benchmarks a binary search where a needle is not found with 100 element
func BenchmarkBinSearch100(b *testing.B) {
	BinSearch(generateArray(100), "no language")
}

// BenchmarkLinSearch1000 benchmarks a binary search where a needle is not found with 1000 element
func BenchmarkBinSearch1000(b *testing.B) {
	BinSearch(generateArray(1000), "no language")
}

// BenchmarkLinSearch10000 benchmarks a binary search where a needle is not found with 10000 element
func BenchmarkBinSearch10000(b *testing.B) {
	BinSearch(generateArray(10000), "no language")
}
