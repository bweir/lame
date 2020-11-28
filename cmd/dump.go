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
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.AddCommand(tokensCmd)
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump compiler outputs.",
}

var tokensCmd = &cobra.Command{
	Use:   "tokens",
	Short: "Dump tokens.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
				"%-3s %-16s (%4d, %4d): %s'%s'\n",
				tok.State[0:3],
				tok.Type,
				tok.Line+1,
				tok.Column+1,
				strings.Repeat("  ", indent),
				tok.Literal,
			)
			if tok.Type == token.ILLEGAL {
				fmt.Printf("Invalid token '%s' (line %d, col %d)\n", tok.Literal, tok.Line, tok.Column)
				os.Exit(1)
			} else if tok.Type == token.UNEXPECTED_EOF {
				fmt.Printf("Unexpected end-of-file (line %d, col %d)\n", tok.Line, tok.Column)
				os.Exit(1)
			}
		}
	},
}
