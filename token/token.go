package token

import "github.com/bweir/lame/token/state"

type Type string

type Token struct {
	Type    Type
	Literal string
	State   state.State
	Line    int
	Column  int
}

const (
	// Special tokens
	NULL   = "NULL"   // do no action
	EOF    = "EOF"    // end of file
	INDENT = "INDENT" // mark indent
	DEDENT = "DEDENT" // mark dedent

	// Literals
	IDENTIFIER = "IDENTIFIER" // fields, table_name

	DECIMAL_NUMBER     = "DECIMAL_NUMBER"
	BINARY_NUMBER      = "BINARY_NUMBER"
	QUATERNARY_NUMBER  = "QUATERNARY_NUMBER"
	HEXADECIMAL_NUMBER = "HEXADECIMAL_NUMBER"

	COMMENT     = "COMMENT"
	DOC_COMMENT = "DOC_COMMENT"
	STRING      = "STRING"
)
