package cmd

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"github.com/lustigo/refingerprint/proc/data"
	"github.com/lustigo/refingerprint/proc/io"
	"github.com/spf13/cobra"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// plotMouseCheckboxCmd represents the plotMouseCheckbox command
var plotMouseCheckboxCmd = &cobra.Command{
	Use:   "plotMouseCheckbox <Single File>",
	Short: "Plots the MousePath from the Rendering of the Captcha to the Checkbox Click",
	Long: `ONLY A SINGLE FILE MUST BE PROVIDED!!
Use without any flags to visualize it in the full size.
Use the width and height flag to scale the visualization.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := GetFiles(args)
		checkErr(err)

		if len(files) > 1 {
			fmt.Println("Please provide only one file to this command.")
			os.Exit(1)
		}

		d, err := io.ParseData(files[0])
		w, h := getScreenSize(cmd, d)

		save, _ := cmd.Flags().GetBool("save")
		file := getFile(save, files)
		saveGraph(w, h, d.MouseCheckbox, d.Screen, d.FramePosition.Captcha, file)
		if !save {
			showWindow(file)
		} else {
			fmt.Printf("The Plot was saved in %v\n", file)
		}
	},
}

// GetFile returns the Filepath to save the plot to
func getFile(toSave bool, files []string) string {
	if !toSave {
		return io.GetTempFile() + ".svg"
	}
	fname := filepath.Base(files[0])
	return filepath.Join(filepath.Dir(files[0]), fname[:len(fname)-len(filepath.Ext(files[0]))]) + ".svg"
}

// GetPlot generates the plot of the Mousepath and saves it to the given file
func getPlot(w, h int16, path []data.NormalizedMouseData, rect data.Rectangle, file string) {
	p, _ := plot.New()
	x := p.X
	x.Min = 0
	x.Max = float64(w)
	y := p.Y
	y.Min = 0
	y.Max = float64(h)
	//p.HideAxes()

	r := plotRectangle(w, h, rect)
	p.Add(r)
	points := plotMousePoints(w, h, path)
	p.Add(points)
	p.Save(vg.Length(w), vg.Length(h), file)
}

// plotMousePoints creates a Scatter Plot from the given MousePath and returns it
func plotMousePoints(w, h int16, path []data.NormalizedMouseData) *plotter.Scatter {
	pts := make(plotter.XYs, len(path))

	for i, point := range path {
		pts[i].X = point.X * float64(w)
		pts[i].Y = (1 - point.Y) * float64(h)
	}

	plotter, _ := plotter.NewScatter(pts)
	return plotter
}

// plotRectangle creates a Polygon Plot from the given Rectangle and returns it
func plotRectangle(w, h int16, rect data.Rectangle) *plotter.Polygon {
	width := float64(w)
	height := float64(h)

	corners := &plotter.XYs{
		{
			X: rect.X * width,
			Y: rect.Y * height,
		},
		{
			X: (rect.X + rect.Width) * width,
			Y: rect.Y * height,
		},
		{
			X: (rect.X + rect.Width) * width,
			Y: (rect.Y - rect.Height) * height,
		},
		{
			X: rect.X * width,
			Y: (rect.Y - rect.Height) * height,
		},
	}

	fmt.Println(rect)
	poly, _ := plotter.NewPolygon(corners)
	poly.Color = color.NRGBA{A: 0x80, R: 0xFF, G: 0x00, B: 0x00}
	return poly
}

// ShowWindows displays the Image in the given file in a GUI
func showWindow(image string) {
	a := app.New()
	w := a.NewWindow("ReFingerPrint MousePathCheckbox")
	i := canvas.NewImageFromFile(image)
	i.FillMode = canvas.ImageFillOriginal
	w.SetContent(i)
	w.ShowAndRun()
}

// Creates and saves the graph of the given path in the given resolution in the given file
func saveGraph(w, h int16, path []data.MouseData, screen data.ScreenInfo, rect data.Rectangle, file string) {
	npath := data.NormalizeMouseData(path, screen)
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

func init() {
	rootCmd.AddCommand(plotMouseCheckboxCmd)

	plotMouseCheckboxCmd.Flags().Int16P("width", "w", -1, "Width of the Screen")
	plotMouseCheckboxCmd.Flags().Int16P("height", "i", -1, "Height of the Screen")
	plotMouseCheckboxCmd.Flags().BoolP("save", "s", false, "Only save the plot, and do not display it")
}
