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
	ILLEGAL = "ILLEGAL"

	EOF     = "EOF"
	SPACE   = "SPACE"
	NEWLINE = "NEWLINE" // \n

	// Literals
	IDENTIFIER  = "IDENTIFIER" // fields, table_name
	NUMBER      = "NUMBER"
	COMMENT     = "COMMENT"
	DOC_COMMENT = "DOC_COMMENT"

	INDENT = "INDENT"
	DEDENT = "DEDENT"
)
