// A lexer implementation of Lame.
//
// 		Hello
package lexer

import (
	"bufio"
	"bytes"
	"container/list"
	"io"

	"github.com/bweir/lame/token"
	"github.com/bweir/lame/token/state"
)

type State int

var eof = rune(0)

type Scanner struct {
	r          *bufio.Reader
	state      state.State
	indent     *list.List
	newIndent  int
	blockStart bool
	line       int
	column     int
	lastChar   rune
	lastColumn int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:      bufio.NewReader(r),
		state:  state.DEFAULT,
		indent: list.New(),
	}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	// fmt.Printf("reading '%s'\n", string(ch))
	if err != nil {
		return eof
	}
	if s.lastChar == '\n' {
		s.lastColumn = s.column
		s.column = 0
		s.line++
	} else {
		s.column++
	}

	s.lastChar = ch
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
	// fmt.Printf("unreading '%s'\n", string(s.lastChar))
	if s.lastChar == '\n' {
		s.line--
		s.column = s.lastColumn
	} else {
		s.column--
	}
}

func (s *Scanner) makeToken(tok token.Type, lit string) token.Token {
	return token.Token{
		Type:    tok,
		Literal: lit,
		State:   s.state,
		Line:    s.line,
		Column:  s.column,
	}
}

func (s *Scanner) Scan() (tok token.Token) {
	currentIndent := 0
	if s.indent.Len() > 0 {
		currentIndent = s.indent.Back().Value.(int)
	}

	if ch := s.read(); ch == eof {
		return s.makeToken(token.EOF, "")
	} else if isNewline(ch) {
		if s.state == state.FUNCTION {
			s.readIndent()
		}
		return s.makeToken(token.NEWLINE, string(ch))
	} else if s.newIndent > currentIndent {
		s.unread()
		return s.scanIndent()
	} else if s.newIndent < currentIndent {
		s.unread()
		return s.scanDedent()
	} else if isSpace(ch) {
		s.unread()
		return s.scanSpace()
	} else if isIdentifier(ch) {
		s.unread()
		return s.scanIdentifier()

	} else if ch == '=' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.EQUAL_TO, "==")
		} else {
			s.unread()
			return s.makeToken(token.ASSIGN, "=")
		}

	} else if ch == '<' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.LESS_THAN_EQUAL_TO, "<=")
		} else if ch == '<' {
			return s.makeToken(token.BITWISE_SHIFT_LEFT, "<<")
		} else {
			s.unread()
			return s.makeToken(token.LESS_THAN, "<")
		}

	} else if ch == '>' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.GREATER_THAN_EQUAL_TO, ">=")
		} else if ch == '>' {
			return s.makeToken(token.BITWISE_SHIFT_RIGHT, ">>")
		} else {
			s.unread()
			return s.makeToken(token.GREATER_THAN, ">")
		}

	} else if ch == '~' {
		if ch = s.read(); ch == '>' {
			return s.makeToken(token.BITWISE_SIGNED_SHIFT_RIGHT, "~>")
		} else {
			s.unread()
			return s.makeToken(token.ILLEGAL, "~")
		}

	} else if ch == '+' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.ADD_ASSIGN, "+=")
		} else {
			s.unread()
			return s.makeToken(token.ADD, "+")
		}

	} else if ch == '-' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.SUBTRACT_ASSIGN, "-=")
		} else {
			s.unread()
			return s.makeToken(token.SUBTRACT, "-")
		}

	} else if ch == '*' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.MULTIPLY_ASSIGN, "*=")
		} else {
			s.unread()
			return s.makeToken(token.MULTIPLY, "*")
		}

	} else if ch == '/' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.DIVIDE_ASSIGN, "/=")
		} else {
			s.unread()
			return s.makeToken(token.DIVIDE, "/")
		}

	} else if ch == '%' {
		if ch = s.read(); ch == '%' {
			return s.scanQuaternaryNumber()
		} else if ch == '=' {
			return s.makeToken(token.MODULO_ASSIGN, "%=")
		} else if !isBinaryDigit(ch) {
			s.unread()
			return s.makeToken(token.MODULO, "%")
		} else {
			s.unread()
			return s.scanBinaryNumber()
		}

	} else if ch == '&' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.BITWISE_AND_ASSIGN, "&=")
		} else {
			s.unread()
			return s.makeToken(token.BITWISE_AND, "&")
		}

	} else if ch == '|' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.BITWISE_OR_ASSIGN, "|=")
		} else {
			s.unread()
			return s.makeToken(token.BITWISE_OR, "|")
		}

	} else if ch == '^' {
		if ch = s.read(); ch == '=' {
			return s.makeToken(token.BITWISE_XOR_ASSIGN, "^=")
		} else {
			s.unread()
			return s.makeToken(token.BITWISE_XOR, "^")
		}

	} else if isDecimalDigit(ch) {
		s.unread()
		return s.scanDecimalNumber()
	} else if ch == '$' {
		return s.scanHexadecimalNumber()

	} else if ch == '"' {
		return s.scanString()

	} else if ch == '.' {
		if ch = s.read(); ch == '.' {
			return s.makeToken(token.RANGE, "..")
		} else {
			s.unread()
			return s.makeToken(token.DOT, ".")
		}

	} else if isLineCommentStart(ch) {
		ch = s.read()
		if ch == eof {
			return s.makeToken(token.EOF, "")
		} else if isLineCommentStart(ch) {
			return s.scanLineDocComment()
		} else {
			s.unread()
			return s.scanLineComment()
		}
	} else if isCommentStart(ch) {
		ch = s.read()
		if ch == eof {
			return s.makeToken(token.EOF, "")
		} else if isCommentStart(ch) {
			return s.scanDocComment()
		} else {
			s.unread()
			return s.scanComment()
		}
	} else {
		s.unread()
		return s.scanOperator()
	}
}

func (s *Scanner) scanSpace() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isSpace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.SPACE, buf.String())
}
