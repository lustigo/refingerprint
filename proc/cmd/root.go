// Package cmd is for the CLI Commands
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "proc",
	Short: "CLI to process ReFingerprint Data",
	Long: `Proc processes data which was collected by the ReFingerprint Webextension.
For example you can check if the datastructure is valid with proc checkIntegrity <File>
Or visualize the mousepath from until the first ReCaptcha Checkbox Click with proc plotMouseCheckpoint <File>

For most of the File-Inputs you can either specify a single file, a glob or a directory.
For some commands that is not possible, but there will be an explicit hint.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
