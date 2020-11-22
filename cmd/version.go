package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of lame",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lame compiler suite 0.0.0")
	},
}
