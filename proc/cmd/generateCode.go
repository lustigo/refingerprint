package cmd

import (
	"github.com/lustigo/refingerprint/proc/gen"
	"github.com/spf13/cobra"
)

// generateCodeCmd represents the generateCode command
var generateCodeCmd = &cobra.Command{
	Use:   "generateCode",
	Short: "Generates code (printed to Stdout) for the data package",
	Long: `Generates code (printed to Stdout) for the data package
The code is being used for the ProcessedFeature struct and the ARFF header.
Also code is generated for  calculating these features.
The features and descriptions are generated for the Movement Features.`,
	Run: func(cmd *cobra.Command, args []string) {
		gen.GenerateCheckbox()
	},
}

func init() {
	rootCmd.AddCommand(generateCodeCmd)
}
