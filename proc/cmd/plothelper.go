package cmd

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"github.com/golang/geo/r2"
	"github.com/lustigo/refingerprint/proc/data"
	"github.com/lustigo/refingerprint/proc/io"
	"github.com/spf13/cobra"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// GetFile returns the Filepath to save the plot to
func getFile(toSave bool, files []string) string {
	if !toSave {
		return io.GetTempFile() + ".svg"
	}
	fname := filepath.Base(files[0])
	return filepath.Join(filepath.Dir(files[0]), fname[:len(fname)-len(filepath.Ext(files[0]))]) + ".svg"
}

// GetPlot generates the plot of the Mousepath and saves it to the given file.
func getPlot(w, h int16, path []data.NormalizedMouseData, rect data.Rectangle, file string) {
	p, _ := plot.New()
	x := p.X
	x.Min = 0
	x.Max = float64(w)
	y := p.Y
	y.Min = 0
	y.Max = float64(h)
	p.HideAxes()

	p.Add(plotAntiCrop(w, h))

	r := plotRectangle(w, h, rect)
	p.Add(r)
	points := plotMousePoints(w, h, path)
	p.Add(points)
	p.Save(vg.Length(w), vg.Length(h), file)
}

// plotAntiCrop returns a Plotter, that draws the whole screen as a transparent rectangle.
func plotAntiCrop(w, h int16) *plotter.Polygon {
	len := r2.Point{X: float64(w), Y: float64(h)}.Norm()
	poly := plotRectangle(w, h, data.Rectangle{X: 0, Y: 0, Width: float64(w) / len, Height: float64(h) / len})
	col := color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
	poly.Color = col
	poly.LineStyle = draw.LineStyle{Color: col}
	return poly
}

// plotMousePoints creates a Scatter Plot from the given MousePath and returns it
func plotMousePoints(w, h int16, path []data.NormalizedMouseData) *plotter.Scatter {
	pts := make(plotter.XYs, len(path))
	len := r2.Point{X: float64(w), Y: float64(h)}.Norm()
	var maxHeight float64
	if len == 0 {
		maxHeight = 0
	} else {
		maxHeight = float64(h) / len
	}

	for i, point := range path {
		pts[i].X = point.X * len
		pts[i].Y = (maxHeight - point.Y) * len
	}

	plotter, _ := plotter.NewScatter(pts)
	return plotter
}

// plotRectangle creates a Polygon Plot from the given Rectangle and returns it
func plotRectangle(w, h int16, rect data.Rectangle) *plotter.Polygon {
	len := r2.Point{X: float64(w), Y: float64(h)}.Norm()
	var maxHeight float64
	if len == 0 {
		maxHeight = 0
	} else {
		maxHeight = float64(h) / len
	}

	corners := &plotter.XYs{
		{
			X: rect.X * len,
			Y: (maxHeight - rect.Y) * len,
		},
		{
			X: (rect.X + rect.Width) * len,
			Y: (maxHeight - rect.Y) * len,
		},
		{
			X: (rect.X + rect.Width) * len,
			Y: (maxHeight - (rect.Y + rect.Height)) * len,
		},
		{
			X: rect.X * len,
			Y: (maxHeight - (rect.Y + rect.Height)) * len,
		},
	}

	poly, _ := plotter.NewPolygon(corners)
	poly.Color = color.NRGBA{A: 0x80, R: 0xFF, G: 0x00, B: 0x00}
	return poly
}

// ShowWindows displays the Image in the given file in a GUI
func showWindow(image string) {
	a := app.New()
	w := a.NewWindow("ReFingerPrint MousePath")
	i := canvas.NewImageFromFile(image)
	i.FillMode = canvas.ImageFillOriginal
	w.SetContent(i)
	w.ShowAndRun()
}

// Creates and saves the graph of the given path in the given resolution in the given file.
func saveGraph(w, h int16, path []data.MouseData, screen data.ScreenInfo, rect data.Rectangle, time data.Time, file string) {
	npath := data.NormalizeMouseData(path, screen, time)
	getPlot(w, h, npath, rect.Normalize(screen), file)
}

// Gets the Screen Size (widht,height) of the Flags (if provided) or from the provided Data
func getScreenSize(cmd *cobra.Command, data *data.Data) (int16, int16) {
	w, err := cmd.Flags().GetInt16("width")
	checkErr(err)
	h, err := cmd.Flags().GetInt16("height")
	checkErr(err)

	if w <= 0 {
		w = int16(data.Screen.Width)
	}
	if h <= 0 {
		h = int16(data.Screen.Height)
	}
	return w, h
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
