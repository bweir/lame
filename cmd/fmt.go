package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fmtCmd)
}

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format Spin code",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("formatting %s ...", args)
	},
}
