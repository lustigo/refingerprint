package cmd

import (
	"fmt"
	"os"

	"github.com/lustigo/refingerprint/proc/io"
	"github.com/spf13/cobra"
)

// plotMouseCheckboxCmd represents the plotMouseCheckbox command
var plotMouseCheckboxCmd = &cobra.Command{
	Use:   "plotMouseCheckbox <Single File>",
	Short: "Plots the MousePath from the Rendering of the Captcha to the Checkbox Click",
	Long: `ONLY A SINGLE FILE MUST BE PROVIDED!!
Use without any flags to visualize it in the full size (RECOMMENDED).
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
		crop, _ := cmd.Flags().GetBool("crop")
		file := getFile(save, files)
		saveGraph(w, h, d.MouseCheckbox, d.Screen, d.FramePosition.Captcha, d.Time, crop, file)
		if !save {
			showWindow(file)
		} else {
			fmt.Printf("The Plot was saved in %v\n", file)
		}
	},
}

func init() {
	rootCmd.AddCommand(plotMouseCheckboxCmd)

	plotMouseCheckboxCmd.Flags().Int16P("width", "w", -1, "Width of the Screen")
	plotMouseCheckboxCmd.Flags().Int16P("height", "i", -1, "Height of the Screen")
	plotMouseCheckboxCmd.Flags().BoolP("crop", "c", false, "Crop the plot to the mousepath section")
	plotMouseCheckboxCmd.Flags().BoolP("save", "s", false, "Only save the plot, and do not display it")
}
