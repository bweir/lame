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
		var tok token.Token

		indent := 0
		for tok.Type != token.EOF {
			tok = scanner.Scan()
			if tok.Type == token.PRI || tok.Type == token.PUB {
				fmt.Printf("\n")
			} else if tok.Type == token.INDENT {
				indent++
			} else if tok.Type == token.DEDENT {
				indent--
			}
			fmt.Printf(
				"%s %12s(%3d, %3d): %s'%s'\n",
				tok.State[0:2],
				tok.Type,
				tok.Line+1,
				tok.Column+1,
				strings.Repeat("  ", indent),
				tok.Literal,
			)
			if tok.Type == token.ILLEGAL {
				fmt.Printf("Invalid token '%s' encountered on line %d, col %d\n", tok.Literal, tok.Line, tok.Column)
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
