package gen

import "testing"

// TestGetPlainTyp tests the getPlainType method
func TestGetPlainType(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected string
	}{
		{"Empty string", "", ""},
		{"Array", "[]float64", "float64"},
		{"Normal Type", "float64", "float64"},
		{"Wrong Array", "float[]", "float[]"},
		{"With a small string, just return it", "a", "a"},
	}

	for _, c := range cases {
		if res := getPlainType(c.s); res != c.expected {
			t.Errorf("%s failed: Expected %s but got %s\n", c.name, c.expected, res)
		}
	}
}

// TestGetIsArray tests the isArray method
func TestIsArray(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected bool
	}{
		{"Empty string", "", false},
		{"Array", "[]float64", true},
		{"Normal Type", "float64", false},
		{"Wrong Array", "float[]", false},
		{"Small string", "a", false},
	}

	for _, c := range cases {
		if res := isArray(c.s); res != c.expected {
			t.Errorf("%s failed: Expected %v but got %v\n", c.name, c.expected, res)
		}
	}
}
