package cmd

import (
	"fmt"
	"os"

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
		displayGraph(w, h, d.MouseCheckbox, d.Screen)
	},
}

// GetPlot generates the plot of the Mousepath
func getPlot(w, h int16, path []data.NormalizedMouseData) string {
	file := io.GetTempFile() + ".svg"
	p, _ := plot.New()
	x := p.X
	x.Min = 0
	x.Max = float64(w)
	y := p.Y
	y.Min = 0
	y.Max = float64(h)
	p.HideAxes()

	pts := make(plotter.XYs, len(path))

	for i, point := range path {
		pts[i].X = point.X * float64(w)
		pts[i].Y = (1 - point.Y) * float64(h)
	}

	plotter, _ := plotter.NewScatter(pts)
	p.Add(plotter)
	p.Save(vg.Length(w), vg.Length(h), file)
	return file
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

// Visualizes the given path in the given resolution
func displayGraph(w, h int16, path []data.MouseData, screen data.ScreenInfo) {
	npath := data.NormalizeMouseData(path, screen)
	tmpImage := getPlot(w, h, npath)
	showWindow(tmpImage)
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
	plotMouseCheckboxCmd.Flags().Int16("height", -1, "Height of the Screen")
}
