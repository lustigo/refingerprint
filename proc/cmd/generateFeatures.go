package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/lustigo/refingerprint/proc/data"
	"github.com/lustigo/refingerprint/proc/io"
	"github.com/spf13/cobra"
)

// FeatureFileName declares the filename for the unseparated FeatureFile without a file extension but with a dot
const FeatureFileName = "features."

// generateFeaturesCmd represents the generateFeatures command
var generateFeaturesCmd = &cobra.Command{
	Use:   "generateFeatures <single file or folder> <options>",
	Short: "Generates the features from the given files.",
	Long: `Generates the features from the given files.
	It saves them either as a CSV (default) or ARFF File (use the -f option).
	If there are multiple datasets the features can be saved for each given dataset (use -s option). The default is that all extracted features will be saved in one file.
	If the features are not separated, the extracted features will be saved in $PWD/features.<format>`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := GetFiles(args)
		checkErr(err)

		sep, _ := cmd.Flags().GetBool("separate")
		format, _ := cmd.Flags().GetString("format")

		if format != "csv" {
			format = "arff"
		}

		generateFeatures(sep, files, format)
	},
}

func init() {
	rootCmd.AddCommand(generateFeaturesCmd)

	generateFeaturesCmd.Flags().StringP("format", "f", "csv", "Format of the feature file, either csv or arff")
	generateFeaturesCmd.Flags().BoolP("separate", "s", false, "If multiple datasets are given, save the features in separate files.")
}

// Generates the Features and saves them
func generateFeatures(sep bool, filepaths []string, format string) {
	if sep {
		generateSeparateFeatures(filepaths, format)
		return
	}
	generateSingleFeatureFile(filepaths, format)
}

// Generates the features and saves them separately
func generateSeparateFeatures(filepaths []string, format string) {
	for _, file := range filepaths {
		d, err := io.ParseData(file)
		if err != nil {
			fmt.Printf("Could not parse %v: %v\n", file, err)
			continue
		}

		features := data.ExtractFeatures(d)
		path := filepath.Join(filepath.Dir(file), io.GetFileName(file)+"."+format)

		if format == "csv" {
			io.WriteCSVFile(path, []*data.ProcessedFeatures{features})
		} else {
			io.WriteARFFFile(path, []*data.ProcessedFeatures{features})
		}
	}
}

// Generates the features and saves it in a single file
func generateSingleFeatureFile(filepaths []string, format string) {
	features := make([]*data.ProcessedFeatures, len(filepaths))

	for i, file := range filepaths {
		d, err := io.ParseData(file)
		if err != nil {
			fmt.Printf("Could not parse %v: %v\n", file, err)
			continue
		}
		feature := data.ExtractFeatures(d)
		features[i] = feature
	}

	path := FeatureFileName + format
	if format == "csv" {
		io.WriteCSVFile(path, features)
	} else {
		io.WriteARFFFile(path, features)
	}
}
