package cmd

import (
	"fmt"

	"github.com/lustigo/refingerprint/proc/io"
	"github.com/spf13/cobra"
)

// checkIntegrityCmd represents the checkIntegrity command
var checkIntegrityCmd = &cobra.Command{
	Use:   "checkIntegrity <File/Files/Dir>",
	Short: "Checks if the Datastructure of the given file(s) is valid",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		files, err := GetFiles(args)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, file := range files {
			checkIntegrity(file)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkIntegrityCmd)
}

// Checks the Integrity of the file. Prints the results to Stdout.
func checkIntegrity(filepath string) {
	_, err := io.ParseData(filepath)
	if err != nil {
		fmt.Printf("%v is not valid: %v\n", filepath, err)
	} else {
		fmt.Printf("%v is valid\n", filepath)
	}
}
