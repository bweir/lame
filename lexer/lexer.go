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

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isNewline(ch rune) bool {
	return ch == '\n'
}

func isIdentifier(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isLineCommentStart(ch rune) bool {
	return ch == '\''
}

func isLineCommentEnd(ch rune) bool {
	return ch == '\n'
}

func isCommentStart(ch rune) bool {
	return ch == '{'
}

func isCommentEnd(ch rune) bool {
	return ch == '}'
}

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
	ch := s.read()
	currentIndent := 0
	if s.indent.Len() > 0 {
		currentIndent = s.indent.Back().Value.(int)
	}
	for {
		if ch == eof {
			return s.makeToken(token.EOF, "")
		} else if isNewline(ch) {
			if s.state == state.FUNCTION {
				if tok := s.readIndent(); tok.Type != token.NULL {
					return tok
				}
			}
		} else if s.newIndent > currentIndent {
			s.unread()
			return s.scanIndent()
		} else if s.newIndent < currentIndent {
			s.unread()
			return s.scanDedent()
		} else if isSpace(ch) {
			s.unread()
			if tok := s.scanSpace(); tok.Type != token.NULL {
				return tok
			}
		} else if isIdentifier(ch) {
			s.unread()
			return s.scanIdentifier()
		} else if isDigit(ch) {
			s.unread()
			return s.scanNumber()
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
		ch = s.read()
	}
}

func (s *Scanner) scanSpace() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			return s.makeToken(token.EOF, "")
		} else if !isSpace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.NULL, "")
}

func (s *Scanner) scanNumber() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.NUMBER, buf.String())
}
