package cmd

import (
	"image/color"
	"testing"

	"github.com/lustigo/refingerprint/proc/data"
	"gonum.org/v1/plot/plotter"
)

// TestPlotMousePoints tests the plotMousePoints function
func TestPlotMousePoints(t *testing.T) {
	var cases = []struct {
		name     string
		path     []data.NormalizedMouseData
		w        int16
		h        int16
		expected plotter.XYs
	}{
		{"no path", []data.NormalizedMouseData{}, 10, 10, plotter.XYs{}},
		{"single point", []data.NormalizedMouseData{{X: 0.1, Y: 0.2}}, 10, 15, plotter.XYs{{X: 1, Y: 12}}}, //Y scale will be inverted
		{"three points", []data.NormalizedMouseData{{X: 0.1, Y: 0.2}, {X: 0.2, Y: 0.4}, {X: 0, Y: 0}}, 10, 15, plotter.XYs{{X: 1, Y: 12}, {X: 2, Y: 9}, {X: 0, Y: 15}}},
		{"three points with w,h=0", []data.NormalizedMouseData{{X: 0.1, Y: 0.2}, {X: 0.2, Y: 0.8}, {X: 0, Y: 0}}, 0, 0, plotter.XYs{{X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}}},
	}

	for _, c := range cases {
		output := plotMousePoints(c.w, c.h, c.path)
		if len(c.expected) != len(output.XYs) {
			t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expected, output.XYs)
		}

		for i, point := range output.XYs {
			if c.expected[i] != point {
				t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expected, output.XYs)
				break
			}
		}
	}
}

// TestPlotRectangle tests the plotRectangle function
func TestPlotRectangle(t *testing.T) {
	red := color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0x80}

	var cases = []struct {
		name        string
		rect        data.Rectangle
		w           int16
		h           int16
		expected    plotter.XYs
		expectedCol color.Color
	}{
		{"square at 0,30%", data.Rectangle{X: 0, Y: 0.3, Width: 0.1, Height: 0.1}, 10, 10, plotter.XYs{{X: 0, Y: 7}, {X: 1, Y: 7}, {X: 1, Y: 6}, {X: 0, Y: 6}}, red},
		{"square at 0,0", data.Rectangle{X: 0, Y: 0, Width: 0.1, Height: 0.1}, 10, 10, plotter.XYs{{X: 0, Y: 10}, {X: 1, Y: 10}, {X: 1, Y: 9}, {X: 0, Y: 9}}, red},
		{"square at 10%,30%", data.Rectangle{X: 0.1, Y: 0.3, Width: 0.1, Height: 0.1}, 10, 10, plotter.XYs{{X: 1, Y: 7}, {X: 2, Y: 7}, {X: 2, Y: 6}, {X: 1, Y: 6}}, red},
		{"rectangle at 0,30% w>h", data.Rectangle{X: 0, Y: 0.3, Width: 0.3, Height: 0.1}, 10, 10, plotter.XYs{{X: 0, Y: 7}, {X: 3, Y: 7}, {X: 3, Y: 6}, {X: 0, Y: 6}}, red},
		{"rectangle at 0,30% h>w", data.Rectangle{X: 0, Y: 0.3, Width: 0.1, Height: 0.3}, 10, 10, plotter.XYs{{X: 0, Y: 7}, {X: 1, Y: 7}, {X: 1, Y: 4}, {X: 0, Y: 4}}, red},
	}

	for _, c := range cases {
		output := plotRectangle(c.w, c.h, c.rect)
		if len(output.XYs) != 1 || len(c.expected) != len(output.XYs[0]) {
			t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expected, output.XYs)
		}

		for _, xyer := range output.XYs {
			for i, point := range xyer {
				if c.expected[i] != point {
					t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expected, xyer)
					break
				}
			}
		}

		if output.Color != c.expectedCol {
			t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expectedCol, output.Color)
		}
	}
}

// TestPlotAntiCrop tests the plotAntiCrop function
func TestPlotAntiCrop(t *testing.T) {
	transparent := color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00}

	var cases = []struct {
		name        string
		w           int16
		h           int16
		expected    plotter.XYs
		expectedCol color.Color
	}{
		{"100x100", 100, 100, plotter.XYs{{X: 0, Y: 100}, {X: 100, Y: 100}, {X: 100, Y: 0}, {X: 0, Y: 0}}, transparent},
		{"1x1", 1, 1, plotter.XYs{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: 0}}, transparent},
		{"50x1", 50, 1, plotter.XYs{{X: 0, Y: 1}, {X: 50, Y: 1}, {X: 50, Y: 0}, {X: 0, Y: 0}}, transparent},
		{"1x50", 1, 50, plotter.XYs{{X: 0, Y: 50}, {X: 1, Y: 50}, {X: 1, Y: 0}, {X: 0, Y: 0}}, transparent},
	}

	for _, c := range cases {
		output := plotAntiCrop(c.w, c.h)
		if len(output.XYs) != 1 || len(c.expected) != len(output.XYs[0]) {
			t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expected, output.XYs)
		}

		for _, xyer := range output.XYs {
			for i, point := range xyer {
				if c.expected[i] != point {
					t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expected, xyer)
					break
				}
			}
		}

		if output.Color != c.expectedCol {
			t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expectedCol, output.Color)
		}
		if output.LineStyle.Color != c.expectedCol {
			t.Errorf("%v failed, expected %v but was %v\n", c.name, c.expectedCol, output.Color)
		}
	}
}
