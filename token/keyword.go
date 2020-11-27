package token

const (
	// Keywords
	// Blocks
	CON = "CON"
	DAT = "DAT"
	OBJ = "OBJ"
	PRI = "PRI"
	PUB = "PUB"
	VAR = "VAR"

	// Constants
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Flow Control
	CASE   = "CASE"
	IF     = "IF"
	ELSEIF = "ELSEIF"
	ELSE   = "ELSE"
	NEXT   = "NEXT"
	QUIT   = "QUIT"
	REPEAT = "REPEAT"
	RETURN = "RETURN"

	// Memory
	BYTE      = "BYTE"
	WORD      = "WORD"
	LONG      = "LONG"
	BYTEFILL  = "BYTEFILL"
	WORDFILL  = "WORDFILL"
	LONGFILL  = "LONGFILL"
	BYTEMOVE  = "BYTEMOVE"
	WORDMOVE  = "WORDMOVE"
	LONGMOVE  = "LONGMOVE"
	LOOKUP    = "LOOKUP"
	LOOKUPZ   = "LOOKUPZ"
	LOOKDOWN  = "LOOKDOWN"
	LOOKDOWNZ = "LOOKDOWNZ"

	// Logical
	NOT = "NOT"
	AND = "AND"
	OR  = "OR"
)
