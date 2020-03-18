/*
Copyright © 2020 Lukas Klassen <luftigo@yahoo.de>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/lustigo/refingerprint/proc/io"
	"github.com/spf13/cobra"
)

// plotMouseRestCmd represents the plotMouseRest command
var plotMouseRestCmd = &cobra.Command{
	Use:   "plotMouseRest <Single File>",
	Short: "Plots the MousePath from the Checkbox Click until the Captcha is solved",
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
		saveGraph(w, h, d.MouseRest, d.Screen, d.FramePosition.Task, d.Time, crop, file)
		if !save {
			showWindow(file)
		} else {
			fmt.Printf("The Plot was saved in %v\n", file)
		}
	},
}

func init() {
	rootCmd.AddCommand(plotMouseRestCmd)

	plotMouseRestCmd.Flags().Int16P("width", "w", -1, "Width of the Screen")
	plotMouseRestCmd.Flags().Int16P("height", "i", -1, "Height of the Screen")
	plotMouseRestCmd.Flags().BoolP("crop", "c", false, "Crop the plot to the mousepath section")
	plotMouseRestCmd.Flags().BoolP("save", "s", false, "Only save the plot, and do not display it")
}
