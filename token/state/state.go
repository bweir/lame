package state

type State string

const (
	DEFAULT     = "DEFAULT"
	COMMENT     = "COMMENT"
	DOC_COMMENT = "DOC_COMMENT"
	FUNCTION    = "FUNCTION"
	VARIABLE    = "VARIABLE"
	DATA        = "DATA"
	OBJECT      = "OBJECT"
	CONSTANT    = "CONSTANT"
)
