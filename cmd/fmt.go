package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bweir/lame/lexer"
	"github.com/bweir/lame/token"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fmtCmd)
}

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format code",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}

		scanner := lexer.NewScanner(strings.NewReader(string(text)))
		var tok token.Token

		indent := 0
		lineStart := true
		for tok.Type != token.EOF {
			tok = scanner.Scan()
			if tok.Type == token.ILLEGAL {
				fmt.Printf("Invalid token '%s' (line %d, col %d)\n", tok.Literal, tok.Line, tok.Column)
				os.Exit(1)
			} else if tok.Type == token.UNEXPECTED_EOF {
				fmt.Printf("Unexpected end-of-file (line %d, col %d)\n", tok.Line, tok.Column)
				os.Exit(1)
			}
			if tok.Type == token.INDENT {
				indent += 4
			} else if tok.Type == token.DEDENT {
				indent -= 4
			} else if tok.Type == token.NEWLINE {
				fmt.Printf("\n")
				lineStart = true
			} else {
				if lineStart {
					fmt.Printf(strings.Repeat(" ", indent))
					lineStart = false
				}
				if tok.Type == token.DOC_COMMENT {
					temp := strings.Split(tok.Literal, "\n")
					for i := 0; i < len(temp); i++ {
						fmt.Printf("'' %s\n", temp[i])
					}
				} else if tok.Type == token.STRING {
					fmt.Printf("\"%s\"", tok.Literal)
				} else if tok.Type == token.INDENT {
					fmt.Printf("      ")
				} else {
					fmt.Printf("%s", tok.Literal)
				}

			}
		}
		fmt.Printf("\n")
	},
}
