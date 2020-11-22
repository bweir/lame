package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/bweir/lame/lexer"
	"github.com/bweir/lame/token"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a Spin object",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("building ...")

		text, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}

		scanner := lexer.NewScanner(strings.NewReader(string(text)))
		// var print_now bool
		var tok token.Token

		for tok.Type != token.EOF {
			tok = scanner.Scan()
			fmt.Printf("%s: '%s'\n", tok.Type, tok.Literal)
			if tok.Type == token.ILLEGAL {
				fmt.Printf("Invalid token: %s '%s'\n", tok.Type, tok.Literal)
				os.Exit(1)
			}
			// if tok == PUB {
			// 	print_now = true
			// }
			// if tok == NEWLINE {
			// 	print_now = false
			// }
			// if tok == DOC_COMMENT {
			// 	fmt.Println()
			// 	fmt.Println(lit)
			// } else if print_now {
			// 	if isBlock(tok) {
			// 		fmt.Printf("### ")
			// 	} else {
			// 		fmt.Printf(lit)
			// 	}
			// }
		}

	},
}
