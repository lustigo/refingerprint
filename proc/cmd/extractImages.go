package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lustigo/refingerprint/proc/io"
	"github.com/spf13/cobra"
	"github.com/vincent-petithory/dataurl"
)

// extractImagesCmd represents the extractImages command
var extractImagesCmd = &cobra.Command{
	Use:   "extractImages",
	Short: "Extracts the images of that file",
	Long: `Extracts all the images of the taks of the given file or files.
	It saves the extracted images as filename_tasknumber_imagenumber.extension .`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := GetFiles(args)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, file := range files {
			extractImages(file)
		}
	},
}

func init() {
	rootCmd.AddCommand(extractImagesCmd)
}

// Extracts the images of the given file
func extractImages(filename string) {
	data, err := io.ParseData(filename)
	if err != nil {
		fmt.Printf("Could not parse %v\n", filename)
		return
	}

	for i, task := range data.TaskEvents {
		for j, img := range task.Images {
			url, err := dataurl.DecodeString(img)
			if err != nil {
				fmt.Printf("Could not parse image %v.%v in %v: %v\n", i, j, filename, err)
			}
			f, err := getFileName(filename, i, j, url)
			if err != nil {
				continue
			}

			_, err = f.Write(url.Data)
			if err != nil {
				fmt.Printf("Could not write to %v: %v\n", f.Name(), err)
			}
			f.Close()
		}
	}
}

func getFileName(filename string, i int, j int, url *dataurl.DataURL) (*os.File, error) {
	imgName := strings.Join([]string{io.GetFileName(filename), "_", strconv.Itoa(i), "_", strconv.Itoa(j), ".", url.MediaType.Subtype}, "")
	imgPathName := filepath.Join(filepath.Dir(filename), imgName)

	var err error
	var f *os.File
	if _, err = os.Stat(imgPathName); err != nil {
		f, err = os.Create(imgPathName)
	} else {
		f, err = os.Open(imgPathName)
	}
	if err != nil {
		fmt.Printf("Could not open %v: %v\n", imgPathName, err)
		return nil, err
	}
	return f, nil

}
