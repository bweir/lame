package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/bweir/lame/token"
)

var eof = rune(0)

func isBlock(tok token.Type) bool {
	switch tok {
	case token.CON:
		fallthrough
	case token.DAT:
		fallthrough
	case token.OBJ:
		fallthrough
	case token.PRI:
		fallthrough
	case token.PUB:
		fallthrough
	case token.VAR:
		return true
	}
	return false
}

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
	r      *bufio.Reader
	line   int
	column int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), line: 0, column: 0}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) makeToken(tok token.Type, lit string) token.Token {
	return token.Token{
		Type:    tok,
		Literal: lit,
		Line:    s.line,
		Column:  s.column,
	}
}

func (s *Scanner) Scan() (tok token.Token) {
	ch := s.read()

	if isSpace(ch) {
		s.unread()
		return s.scanSpace()
	} else if isIdentifier(ch) {
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

	switch ch {
	case eof:
		return s.makeToken(token.EOF, "")
	case '\n':
		return s.makeToken(token.NEWLINE, string(ch))
	case ',':
		return s.makeToken(token.COMMA, string(ch))
	case '=':
		return s.makeToken(token.EQUAL, string(ch))

	case '<':
		return s.makeToken(token.LESS_THAN, string(ch))
	case '>':
		return s.makeToken(token.GREATER_THAN, string(ch))
	case '~':
		return s.makeToken(token.TILDE, string(ch))
	case '&':
		return s.makeToken(token.AMPERSAND, string(ch))

	case ':':
		return s.makeToken(token.COLON, string(ch))
	case '"':
		return s.makeToken(token.QUOTE_DOUBLE, string(ch))
	case '.':
		return s.makeToken(token.DOT, string(ch))
	case '|':
		return s.makeToken(token.PIPE, string(ch))
	case '!':
		return s.makeToken(token.EXCLAMATION_MARK, string(ch))
	case '@':
		return s.makeToken(token.AT, string(ch))
	case '#':
		return s.makeToken(token.POUND, string(ch))
	case '$':
		return s.makeToken(token.DOLLAR, string(ch))
	case '%':
		return s.makeToken(token.PERCENT, string(ch))

	case '+':
		return s.makeToken(token.PLUS, string(ch))
	case '-':
		return s.makeToken(token.MINUS, string(ch))
	case '/':
		return s.makeToken(token.SLASH, string(ch))
	case '*':
		return s.makeToken(token.ASTERISK, string(ch))

	case '(':
		return s.makeToken(token.PAREN_OPEN, string(ch))
	case ')':
		return s.makeToken(token.PAREN_CLOSE, string(ch))
	case '[':
		return s.makeToken(token.BRACKET_OPEN, string(ch))
	case ']':
		return s.makeToken(token.BRACKET_CLOSE, string(ch))
	}

	return s.makeToken(token.ILLEGAL, string(ch))
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

	return s.makeToken(token.WS, buf.String())
}

func (s *Scanner) scanLineDocComment() (tok token.Token) {
	tok = s.scanLineComment()
	tok.Type = token.DOC_COMMENT
	return tok
}

func (s *Scanner) scanLineComment() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof || isLineCommentEnd(ch) {
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.COMMENT, buf.String())
}

func (s *Scanner) scanDocComment() (tok token.Token) {
	tok = s.scanComment()
	if ch := s.read(); !isCommentEnd(ch) {
		return s.makeToken(token.ILLEGAL, string(ch))
	}
	tok.Type = token.DOC_COMMENT
	return tok
}

func (s *Scanner) scanComment() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof || isCommentEnd(ch) {
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.COMMENT, buf.String())
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

func (s *Scanner) scanIdentifier() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isIdentifier(ch) && !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {

	// Blocks
	case "CON":
		return s.makeToken(token.CON, buf.String())
	case "DAT":
		return s.makeToken(token.DAT, buf.String())
	case "OBJ":
		return s.makeToken(token.OBJ, buf.String())
	case "PRI":
		return s.makeToken(token.PRI, buf.String())
	case "PUB":
		return s.makeToken(token.PUB, buf.String())
	case "VAR":
		return s.makeToken(token.VAR, buf.String())

	// Constants
	case "TRUE":
		return s.makeToken(token.TRUE, buf.String())
	case "FALSE":
		return s.makeToken(token.FALSE, buf.String())
	case "POSX":
		return s.makeToken(token.POSX, buf.String())
	case "NEGX":
		return s.makeToken(token.NEGX, buf.String())
	case "PI":
		return s.makeToken(token.PI, buf.String())
	case "RCFAST":
		return s.makeToken(token.RCFAST, buf.String())
	case "RCSLOW":
		return s.makeToken(token.RCSLOW, buf.String())
	case "XINPUT":
		return s.makeToken(token.XINPUT, buf.String())
	case "XTAL1":
		return s.makeToken(token.XTAL1, buf.String())
	case "XTAL2":
		return s.makeToken(token.XTAL2, buf.String())
	case "XTAL3":
		return s.makeToken(token.XTAL3, buf.String())
	case "PLL1X":
		return s.makeToken(token.PLL1X, buf.String())
	case "PLL2X":
		return s.makeToken(token.PLL2X, buf.String())
	case "PLL4X":
		return s.makeToken(token.PLL4X, buf.String())
	case "PLL8X":
		return s.makeToken(token.PLL8X, buf.String())
	case "PLL16X":
		return s.makeToken(token.PLL16X, buf.String())

	// Variables
	case "RESULT":
		return s.makeToken(token.RESULT, buf.String())

	// Flow Control
	case "ABORT":
		return s.makeToken(token.ABORT, buf.String())
	case "CASE":
		return s.makeToken(token.CASE, buf.String())
	case "IF":
		return s.makeToken(token.IF, buf.String())
	case "IFNOT":
		return s.makeToken(token.IFNOT, buf.String())
	case "NEXT":
		return s.makeToken(token.NEXT, buf.String())
	case "QUIT":
		return s.makeToken(token.QUIT, buf.String())
	case "REPEAT":
		return s.makeToken(token.REPEAT, buf.String())
	case "RETURN":
		return s.makeToken(token.RETURN, buf.String())

	// Memory
	case "BYTE":
		return s.makeToken(token.BYTE, buf.String())
	case "WORD":
		return s.makeToken(token.WORD, buf.String())
	case "LONG":
		return s.makeToken(token.LONG, buf.String())
	case "BYTEFILL":
		return s.makeToken(token.BYTEFILL, buf.String())
	case "WORDFILL":
		return s.makeToken(token.WORDFILL, buf.String())
	case "LONGFILL":
		return s.makeToken(token.LONGFILL, buf.String())
	case "BYTEMOVE":
		return s.makeToken(token.BYTEMOVE, buf.String())
	case "WORDMOVE":
		return s.makeToken(token.WORDMOVE, buf.String())
	case "LONGMOVE":
		return s.makeToken(token.LONGMOVE, buf.String())
	case "LOOKUP":
		return s.makeToken(token.LOOKUP, buf.String())
	case "LOOKUPZ":
		return s.makeToken(token.LOOKUPZ, buf.String())
	case "LOOKDOWN":
		return s.makeToken(token.LOOKDOWN, buf.String())
	case "LOOKDOWNZ":
		return s.makeToken(token.LOOKDOWNZ, buf.String())
	case "STRSIZE":
		return s.makeToken(token.STRSIZE, buf.String())
	case "STRCOMP":
		return s.makeToken(token.STRCOMP, buf.String())

	// Directives
	case "STRING":
		return s.makeToken(token.STRING, buf.String())
	case "CONSTANT":
		return s.makeToken(token.CONSTANT, buf.String())
	case "FLOAT":
		return s.makeToken(token.FLOAT, buf.String())
	case "ROUND":
		return s.makeToken(token.ROUND, buf.String())
	case "TRUNC":
		return s.makeToken(token.TRUNC, buf.String())
	case "FILE":
		return s.makeToken(token.FILE, buf.String())

	// Locks
	case "LOCKNEW":
		return s.makeToken(token.LOCKNEW, buf.String())
	case "LOCKRET":
		return s.makeToken(token.LOCKRET, buf.String())
	case "LOCKCLR":
		return s.makeToken(token.LOCKCLR, buf.String())
	case "LOCKSET":
		return s.makeToken(token.LOCKSET, buf.String())

	// Chip Config
	case "CHIPVER":
		return s.makeToken(token.CHIPVER, buf.String())
	case "CLKMODE":
		return s.makeToken(token.CLKMODE, buf.String())
	case "CLKFREQ":
		return s.makeToken(token.CLKFREQ, buf.String())
	case "CLKSET":
		return s.makeToken(token.CLKSET, buf.String())
	case "_CLKMODE":
		return s.makeToken(token.P_CLKMODE, buf.String())
	case "_CLKFREQ":
		return s.makeToken(token.P_CLKFREQ, buf.String())
	case "_XINFREQ":
		return s.makeToken(token.P_XINFREQ, buf.String())
	case "_STACK":
		return s.makeToken(token.P_STACK, buf.String())
	case "_FREE":
		return s.makeToken(token.P_FREE, buf.String())

	// Registers
	case "CNT":
		return s.makeToken(token.CNT, buf.String())
	case "CTRA":
		return s.makeToken(token.CTRA, buf.String())
	case "CTRB":
		return s.makeToken(token.CTRB, buf.String())
	case "DIRA":
		return s.makeToken(token.DIRA, buf.String())
	case "DIRB":
		return s.makeToken(token.DIRB, buf.String())
	case "INA":
		return s.makeToken(token.INA, buf.String())
	case "INB":
		return s.makeToken(token.INB, buf.String())
	case "OUTA":
		return s.makeToken(token.OUTA, buf.String())
	case "OUTB":
		return s.makeToken(token.OUTB, buf.String())
	case "FRQA":
		return s.makeToken(token.FRQA, buf.String())
	case "FRQB":
		return s.makeToken(token.FRQB, buf.String())
	case "PHSA":
		return s.makeToken(token.PHSA, buf.String())
	case "PHSB":
		return s.makeToken(token.PHSB, buf.String())
	case "VCFG":
		return s.makeToken(token.VCFG, buf.String())
	case "VSCL":
		return s.makeToken(token.VSCL, buf.String())
	case "PAR":
		return s.makeToken(token.PAR, buf.String())
	case "SPR":
		return s.makeToken(token.SPR, buf.String())

	// Process Control
	case "WAITCNT":
		return s.makeToken(token.WAITCNT, buf.String())
	case "WAITPEQ":
		return s.makeToken(token.WAITPEQ, buf.String())
	case "WAITPNE":
		return s.makeToken(token.WAITPNE, buf.String())
	case "WAITVID":
		return s.makeToken(token.WAITVID, buf.String())

	// Cog Control
	case "COGID":
		return s.makeToken(token.COGID, buf.String())
	case "COGNEW":
		return s.makeToken(token.COGNEW, buf.String())
	case "COGINIT":
		return s.makeToken(token.COGINIT, buf.String())
	case "COGSTOP":
		return s.makeToken(token.COGSTOP, buf.String())
	case "REBOOT":
		return s.makeToken(token.REBOOT, buf.String())
	}

	// Otherwise return s.makeToken(as a regular identifier.
	return s.makeToken(token.IDENTIFIER, buf.String())
}
