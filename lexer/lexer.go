package lexer

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
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
	r           *bufio.Reader
	state       state.State
	indent      *list.List
	line_start  bool
	line        int
	column      int
	last_char   rune
	last_column int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:           bufio.NewReader(r),
		state:       state.DEFAULT,
		indent:      list.New(),
		line_start:  false,
		line:        0,
		column:      0,
		last_char:   0,
		last_column: 0,
	}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if s.last_char == '\n' {
		s.last_column = s.column
		s.column = 0
		s.line++
		if s.state == state.FUNCTION {
			s.line_start = true
		}
	} else {
		s.column++
	}

	s.last_char = ch
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
	if s.last_char == '\n' {
		s.line--
		s.column = s.last_column
		if s.state == state.FUNCTION {
			s.line_start = false
		}
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

	if isSpace(ch) {
		if s.line_start {
			s.unread()
			s.line_start = false
			return s.scanIndent()
		} else {
			s.unread()
			return s.scanSpace()
		}
	}

	s.line_start = false

	if isIdentifier(ch) {
		s.unread()
		return s.scanIdentifier()
	} else if isDigit(ch) {
		s.unread()
		return s.scanNumber()
	} else if isLineCommentStart(ch) {
		ch := s.read()
		if ch == eof {
			return s.makeToken(token.EOF, "")
		} else if isLineCommentStart(ch) {
			return s.scanLineDocComment()
		} else {
			s.unread()
			return s.scanLineComment()
		}
	} else if isCommentStart(ch) {
		ch := s.read()
		if ch == eof {
			return s.makeToken(token.EOF, "")
		} else if isCommentStart(ch) {
			return s.scanDocComment()
		} else {
			s.unread()
			return s.scanComment()
		}
	}

	s.unread()
	return s.scanOperator()
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

func (s *Scanner) scanIndent() (tok token.Token) {
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

	for e := s.indent.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}

	// add up indents
	current_indent := 0
	for e := s.indent.Front(); e != nil; e = e.Next() {
		current_indent += e.Value.(int)
	}

	new_indent := len(buf.String())
	// total += len(buf.String())

	if new_indent > current_indent {
		// indenting
		s.indent.PushBack(new_indent)
		return s.makeToken(token.INDENT, "-->")
	} else if new_indent < current_indent {
		// dedenting
		s.indent.Remove(s.indent.Back())
		return s.makeToken(token.DEDENT, "<--")
	}
	return s.makeToken(token.ILLEGAL, buf.String())
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
