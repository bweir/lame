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
	ILLEGAL = "ILLEGAL" // invalid character
	NULL    = "NULL"    // do no action
	EOF     = "EOF"     // end of file

	// Literals
	IDENTIFIER  = "IDENTIFIER" // fields, table_name
	NUMBER      = "NUMBER"
	COMMENT     = "COMMENT"
	DOC_COMMENT = "DOC_COMMENT"

	INDENT = "INDENT"
	DEDENT = "DEDENT"
)
