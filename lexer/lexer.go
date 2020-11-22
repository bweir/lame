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
	r             *bufio.Reader
	state         state.State
	indent        *list.List
	initialIndent int
	blockStart    bool
	lineStart     bool
	line          int
	column        int
	lastChar      rune
	lastColumn    int
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
	for {
		if ch == eof {
			return s.makeToken(token.EOF, "")
		} else if isNewline(ch) {
			// fmt.Println("go newline", ch)
			//s.lineStart = false
			//} else if s.lineStart {
			//	s.unread()
			//	s.scanSpace()
			// if s.state == state.FUNCTION {
			// 	tok = s.scanIndentLevel()
			// 	if tok.Type != token.NULL {
			// 		return tok
			// 	}
			// }
		} else if isSpace(ch) {
			// fmt.Println("go space", ch)
			s.unread()
			s.scanSpace()
		} else if isIdentifier(ch) {
			// fmt.Println("go identifier", ch)
			s.unread()
			return s.scanIdentifier()
		} else if isDigit(ch) {
			// fmt.Println("go digit", ch)
			s.unread()
			return s.scanNumber()
		} else if isLineCommentStart(ch) {
			// fmt.Println("go line comment", ch)
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
			// fmt.Println("go comment", ch)
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
			// fmt.Println("go operator", ch)
			s.unread()
			return s.scanOperator()
		}
		ch = s.read()
	}
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
