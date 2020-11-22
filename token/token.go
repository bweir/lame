package token

type Type string

type Token struct {
	Type    Type
	Literal string
	Line    int
	Column  int
}

const (
	// Special tokens
	ILLEGAL = "ILLEGAL"

	EOF     = "EOF"
	WS      = "WS"
	NEWLINE = "NEWLINE" // \n

	// Literals
	IDENTIFIER  = "IDENTIFIER" // fields, table_name
	NUMBER      = "NUMBER"
	COMMENT     = "COMMENT"
	DOC_COMMENT = "DOC_COMMENT"

	INDENT = "INDENT"
	DEDENT = "DEDENT"
)
