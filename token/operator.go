package token

const (
	// Operators
	ADD      = "ADD"      // +
	SUBTRACT = "SUBTRACT" // -
	MULTIPLY = "MULTIPLY" // *
	DIVIDE   = "DIVIDE"   // /
	MODULO   = "MODULO"   // %
	ASSIGN   = "ASSIGN"   // =

	ADD_ASSIGN      = "ADD_ASSIGN"      // +=
	SUBTRACT_ASSIGN = "SUBTRACT_ASSIGN" // -=
	MULTIPLY_ASSIGN = "MULTIPLY_ASSIGN" // *=
	DIVIDE_ASSIGN   = "DIVIDE_ASSIGN"   // /=
	MODULO_ASSIGN   = "MODULO_ASSIGN"   // /=

	AT    = "AT"    // @
	POUND = "POUND" // #

	DOT   = "DOT"   // .
	RANGE = "RANGE" // ..
	PIPE  = "PIPE"  // |

	EQUAL_TO              = "EQUAL_TO"              // ==
	LESS_THAN             = "LESS_THAN"             // <
	GREATER_THAN          = "GREATER_THAN"          // >
	LESS_THAN_EQUAL_TO    = "LESS_THAN_EQUAL_TO"    // <=
	GREATER_THAN_EQUAL_TO = "GREATER_THAN_EQUAL_TO" // >=

	BITWISE_AND        = "BITWISE_AND"        // &
	BITWISE_AND_ASSIGN = "BITWISE_AND_ASSIGN" // &=
	BITWISE_OR         = "BITWISE_OR"         // |
	BITWISE_OR_ASSIGN  = "BITWISE_OR_ASSIGN"  // |=
	BITWISE_XOR        = "BITWISE_XOR"        // ^
	BITWISE_XOR_ASSIGN = "BITWISE_XOR_ASSIGN" // ^=
	BITWISE_NOT        = "BITWISE_NOT"        // !

	BITWISE_SHIFT_LEFT         = "BITWISE_SHIFT_LEFT"         // <<
	BITWISE_SHIFT_RIGHT        = "BITWISE_SHIFT_RIGHT"        // >>
	BITWISE_ROTATE_LEFT        = "BITWISE_ROTATE_LEFT"        // <-
	BITWISE_ROTATE_RIGHT       = "BITWISE_ROTATE_RIGHT"       // ->
	BITWISE_REVERSE            = "BITWISE_REVERSE"            // ><
	BITWISE_SIGNED_SHIFT_RIGHT = "BITWISE_SIGNED_SHIFT_RIGHT" // ~>
	BITWISE_SIGN_EXTEND_7      = "BITWISE_SIGN_EXTEND_7"      // ~
	BITWISE_SIGN_EXTEND_15     = "BITWISE_SIGN_EXTEND_15"     // ~~

	// Misc characters
	BRACKET_OPEN  = "BRACKET_OPEN"  // [
	BRACKET_CLOSE = "BRACKET_CLOSE" // ]
	COMMA         = "COMMA"         // ,
	PAREN_OPEN    = "PAREN_OPEN"    // (
	PAREN_CLOSE   = "PAREN_CLOSE"   // )
	COLON         = "COLON"         // :
	BRACE_OPEN    = "BRACE_OPEN"    // {
	BRACE_CLOSE   = "BRACE_CLOSE"   // }
)
