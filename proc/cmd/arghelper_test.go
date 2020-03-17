package cmd

import (
	"testing"
)

// TestContains tests the Contains Function
func TestContains(t *testing.T) {
	var cases = []struct {
		name     string
		needle   string
		haystack []string
		expected bool
	}{
		{"noHaystack with needle", "needle", []string{}, false},
		{"only correct item in haystack", "needle", []string{"needle"}, true},
		{"only false item in haystack", "needle", []string{"notneedle"}, false},
		{"correct item last position", "needle", []string{"notneedle", "alsonot", "nope", "needle"}, true},
		{"correct item first position", "needle", []string{"needle", "notneedle", "alsonot", "nope"}, true},
		{"correct item middle position", "needle", []string{"notneedle", "alsonot", "needle", "nope", "needlexx"}, true},
		{"not correct", "needle", []string{"notneedle", "alsonot", "nope", "needlexxx"}, false},
		{"empty needle with haystack", "", []string{"notneedle", "alsonot", "nope", "needle"}, false},
		{"empty needle without haystack", "", []string{}, false},
	}

	for _, c := range cases {
		if Contains(c.haystack, c.needle) != c.expected {
			t.Errorf("%v returns %v\n", c.name, !c.expected)
		}
	}
}

//TestCheckFileAndAppend tests the CheckFileAndAppend Function
func TestCheckFileAndAppend(t *testing.T) {
	var cases = []struct {
		name     string
		file     string
		paths    []string
		expected []string
	}{
		{"correct txt file and empty path", "this.txt", []string{}, []string{"this.txt"}},
		{"correct json file and empty path", "this.json", []string{}, []string{"this.json"}},
		{"incorrect file and empty path", "lifeofensidias.de", []string{}, []string{}},
		{"incorrect txt file and empty path", "this.txt.zip", []string{}, []string{}},
		{"incorrect json file and empty path", "this.json.mp4", []string{}, []string{}},
		{"correct txt file and non-empty path", "this.txt", []string{"hallo"}, []string{"hallo", "this.txt"}},
		{"correct json file and non-empty path", "this.json", []string{"hey"}, []string{"hey", "this.json"}},
		{"incorrect file and non-empty path", "lifeofensidias.de", []string{"ola"}, []string{"ola"}},
		{"incorrect txt file and non-empty path", "this.txt.zip", []string{"moin"}, []string{"moin"}},
		{"incorrect json file and non-empty path", "this.json.mp4", []string{"gutn"}, []string{"gutn"}},
	}

	for _, c := range cases {
		output := checkFileAndAppend(c.file, c.paths)
		if len(output) != len(c.expected) {
			t.Errorf("%v returns %v, but %v was expected\n", c.name, output, c.expected)
		}

		for i, o := range output {
			if o != c.expected[i] {
				t.Errorf("%v returns %v, but %v was expected\n", c.name, output, c.expected)
				break
			}
		}
	}
}
