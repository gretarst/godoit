package godoit

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godoit",
	Short: "godoit - todo list manager.",
	Long:  "godoit is a command line application to manage todo lists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error running the CLI.")
		os.Exit(1)
	}
}
