package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"./lexer"
)

var filename string

func main() {
	// flag.StringVar(&filename, "filename", "", "Spin file to parse")
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	text, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	scanner := lexer.NewScanner(strings.NewReader(string(text)))
	// var print_now bool
	var tok lexer.Token
	var lit string

	for tok != lexer.EOF {
		tok, lit = scanner.Scan()
		fmt.Println(tok, lit)
		if tok == lexer.ILLEGAL {
			fmt.Printf("Invalid token: %q\n", lit)
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
}
