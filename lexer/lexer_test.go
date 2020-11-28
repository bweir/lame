package lexer_test

import (
	"strings"
	"testing"

	"github.com/bweir/lame/lexer"
	"github.com/bweir/lame/token"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		src     string
		Type    token.Type
		Literal string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{src: ``, Type: token.EOF},
		{src: ` `, Type: token.SPACE, Literal: ` `},
		{src: "\t", Type: token.SPACE, Literal: "\t"},
		{src: "\n", Type: token.NEWLINE, Literal: "\n"},

		// Operators
		// Arithmetic
		{src: `+`, Type: token.ADD, Literal: `+`},
		{src: `+=`, Type: token.ADD_ASSIGN, Literal: `+=`},
		{src: `-`, Type: token.SUBTRACT, Literal: `-`},
		{src: `-=`, Type: token.SUBTRACT_ASSIGN, Literal: `-=`},
		{src: `*`, Type: token.MULTIPLY, Literal: `*`},
		{src: `*=`, Type: token.MULTIPLY_ASSIGN, Literal: `*=`},
		{src: `/`, Type: token.DIVIDE, Literal: `/`},
		{src: `/=`, Type: token.DIVIDE_ASSIGN, Literal: `/=`},
		{src: `%`, Type: token.MODULO, Literal: `%`},
		{src: `%=`, Type: token.MODULO_ASSIGN, Literal: `%=`},

		// Bitwise
		{src: `&`, Type: token.BITWISE_AND, Literal: `&`},
		{src: `&=`, Type: token.BITWISE_AND_ASSIGN, Literal: `&=`},
		{src: `|`, Type: token.BITWISE_OR, Literal: `|`},
		{src: `|=`, Type: token.BITWISE_OR_ASSIGN, Literal: `|=`},
		{src: `^`, Type: token.BITWISE_XOR, Literal: `^`},
		{src: `^=`, Type: token.BITWISE_XOR_ASSIGN, Literal: `^=`},
		{src: `!`, Type: token.BITWISE_NOT, Literal: `!`},

		{src: `<<`, Type: token.BITWISE_SHIFT_LEFT, Literal: `<<`},
		{src: `>>`, Type: token.BITWISE_SHIFT_RIGHT, Literal: `>>`},
		{src: `~>`, Type: token.BITWISE_SIGNED_SHIFT_RIGHT, Literal: `~>`},

		{src: `~`, Type: token.ILLEGAL, Literal: `~`},

		// Comparison
		{src: `=`, Type: token.ASSIGN, Literal: `=`},
		{src: `==`, Type: token.EQUAL_TO, Literal: `==`},
		{src: `<`, Type: token.LESS_THAN, Literal: `<`},
		{src: `<=`, Type: token.LESS_THAN_EQUAL_TO, Literal: `<=`},
		{src: `>`, Type: token.GREATER_THAN, Literal: `>`},
		{src: `>=`, Type: token.GREATER_THAN_EQUAL_TO, Literal: `>=`},
		{src: `.`, Type: token.DOT, Literal: `.`},
		{src: `..`, Type: token.RANGE, Literal: `..`},

		// Numbers

		// Decimal numbers
		{src: `36564`, Type: token.DECIMAL_NUMBER, Literal: `36564`},
		{src: `36_564`, Type: token.DECIMAL_NUMBER, Literal: `36_564`}, // group separators work
		{src: `036`, Type: token.DECIMAL_NUMBER, Literal: `036`},       // zero padding is fine
		{src: `036AF`, Type: token.DECIMAL_NUMBER, Literal: `036`},     // error must be caught by parser

		// Binary numbers
		{src: `101010`, Type: token.DECIMAL_NUMBER, Literal: `101010`}, // still decimal here
		{src: `%101010`, Type: token.BINARY_NUMBER, Literal: `101010`}, // now binary
		{src: `%101_010`, Type: token.BINARY_NUMBER, Literal: `101_010`},
		{src: `%10123`, Type: token.BINARY_NUMBER, Literal: `101`}, // error must be caught in parser

		// Quaternary numbers
		{src: `10123`, Type: token.DECIMAL_NUMBER, Literal: `10123`},      // still decimal here
		{src: `%%10123`, Type: token.QUATERNARY_NUMBER, Literal: `10123`}, // now quaternary
		{src: `%%101_010`, Type: token.QUATERNARY_NUMBER, Literal: `101_010`},
		{src: `%%10179`, Type: token.QUATERNARY_NUMBER, Literal: `101`}, // error must be caught in parser

		// Hexadecimal numbers
		{src: `1ACD3`, Type: token.DECIMAL_NUMBER, Literal: `1`},
		{src: `ACD3`, Type: token.IDENTIFIER, Literal: `ACD3`},
		{src: `$1ACD3`, Type: token.HEXADECIMAL_NUMBER, Literal: `1ACD3`},
		{src: `$1acd3`, Type: token.HEXADECIMAL_NUMBER, Literal: `1acd3`},
		{src: `$1ac_d3`, Type: token.HEXADECIMAL_NUMBER, Literal: `1ac_d3`},

		// Identifiers
		{src: `foobar`, Type: token.IDENTIFIER, Literal: `foobar`},
		{src: `foo_bar`, Type: token.IDENTIFIER, Literal: `foo_bar`},
		{src: `__foo__`, Type: token.IDENTIFIER, Literal: `__foo__`},
		{src: `_36`, Type: token.IDENTIFIER, Literal: `_36`}, // this is valid

		// Keywords
		// Blocks
		{src: `con`, Type: token.CON, Literal: `con`},
		{src: `dat`, Type: token.DAT, Literal: `dat`},
		{src: `obj`, Type: token.OBJ, Literal: `obj`},
		{src: `pri`, Type: token.PRI, Literal: `pri`},
		{src: `pub`, Type: token.PUB, Literal: `pub`},
		{src: `var`, Type: token.VAR, Literal: `var`},

		{src: `PUB`, Type: token.PUB, Literal: `PUB`},

		// Constants
		{src: `true`, Type: token.TRUE, Literal: `true`},
		{src: `false`, Type: token.FALSE, Literal: `false`},

		// Flow Control
		{src: `case`, Type: token.CASE, Literal: `case`},
		{src: `if`, Type: token.IF, Literal: `if`},
		{src: `elseif`, Type: token.ELSEIF, Literal: `elseif`},
		{src: `else`, Type: token.ELSE, Literal: `else`},
		{src: `next`, Type: token.NEXT, Literal: `next`},
		{src: `quit`, Type: token.QUIT, Literal: `quit`},
		{src: `repeat`, Type: token.REPEAT, Literal: `repeat`},
		{src: `from`, Type: token.FROM, Literal: `from`},
		{src: `to`, Type: token.TO, Literal: `to`},
		{src: `step`, Type: token.STEP, Literal: `step`},
		{src: `while`, Type: token.WHILE, Literal: `while`},
		{src: `until`, Type: token.UNTIL, Literal: `until`},
		{src: `return`, Type: token.RETURN, Literal: `return`},

		// Memory
		{src: `byte`, Type: token.BYTE, Literal: `byte`},
		{src: `word`, Type: token.WORD, Literal: `word`},
		{src: `long`, Type: token.LONG, Literal: `long`},

		// Logical
		{src: `not`, Type: token.NOT, Literal: `not`},
		{src: `and`, Type: token.AND, Literal: `and`},
		{src: `or`, Type: token.OR, Literal: `or`},
	}

	for i, tt := range tests {
		s := lexer.NewScanner(strings.NewReader(tt.src))
		tok := s.Scan()
		if tt.Type != tok.Type {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.src, tt.Type, tok.Type, tok.Literal)
		} else if tt.Literal != tok.Literal {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.src, tt.Literal, tok.Literal)
		}
	}
}
