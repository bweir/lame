package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

func isBlock(tok Token) bool {
	return tok >= CON && tok <= VAR
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
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
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

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) Scan() (tok Token, lit string) {
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
			return EOF, ""
		} else if isLineCommentStart(ch) {
			return s.scanLineDocComment()
		} else {
			s.unread()
			return s.scanLineComment()
		}
	} else if isCommentStart(ch) {
		ch := s.read()
		if ch == eof {
			return EOF, ""
		} else if isCommentStart(ch) {
			return s.scanDocComment()
		} else {
			s.unread()
			return s.scanComment()
		}
	}

	switch ch {
	case eof:
		return EOF, ""
	case '\n':
		return NEWLINE, string(ch)
	case ',':
		return COMMA, string(ch)
	case '=':
		return EQUAL, string(ch)

	case '<':
		return LESS_THAN, string(ch)
	case '>':
		return GREATER_THAN, string(ch)
	case '~':
		return TILDE, string(ch)
	case '&':
		return AMPERSAND, string(ch)

	case ':':
		return COLON, string(ch)
	case '"':
		return QUOTE_DOUBLE, string(ch)
	case '.':
		return DOT, string(ch)
	case '|':
		return PIPE, string(ch)
	case '!':
		return EXCLAMATION_MARK, string(ch)
	case '@':
		return AT, string(ch)
	case '#':
		return POUND, string(ch)
	case '$':
		return DOLLAR, string(ch)
	case '%':
		return PERCENT, string(ch)

	case '+':
		return PLUS, string(ch)
	case '-':
		return MINUS, string(ch)
	case '/':
		return SLASH, string(ch)
	case '*':
		return ASTERISK, string(ch)

	case '(':
		return PAREN_OPEN, string(ch)
	case ')':
		return PAREN_CLOSE, string(ch)
	case '[':
		return BRACKET_OPEN, string(ch)
	case ']':
		return BRACKET_CLOSE, string(ch)
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) scanSpace() (tok Token, lit string) {
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

	return WS, buf.String()
}

func (s *Scanner) scanLineDocComment() (tok Token, lit string) {
	_, lit = s.scanLineComment()
	return DOC_COMMENT, lit
}

func (s *Scanner) scanLineComment() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof || isLineCommentEnd(ch) {
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return COMMENT, buf.String()
}

func (s *Scanner) scanDocComment() (tok Token, lit string) {
	_, lit = s.scanComment()
	if ch := s.read(); !isCommentEnd(ch) {
		return ILLEGAL, string(ch)
	}
	return DOC_COMMENT, lit
}

func (s *Scanner) scanComment() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof || isCommentEnd(ch) {
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return COMMENT, buf.String()
}

func (s *Scanner) scanNumber() (tok Token, lit string) {
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

	return NUMBER, buf.String()
}

func (s *Scanner) scanIdentifier() (tok Token, lit string) {
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
		return CON, buf.String()
	case "DAT":
		return DAT, buf.String()
	case "OBJ":
		return OBJ, buf.String()
	case "PRI":
		return PRI, buf.String()
	case "PUB":
		return PUB, buf.String()
	case "VAR":
		return VAR, buf.String()

	// Constants
	case "TRUE":
		return TRUE, buf.String()
	case "FALSE":
		return FALSE, buf.String()
	case "POSX":
		return POSX, buf.String()
	case "NEGX":
		return NEGX, buf.String()
	case "PI":
		return PI, buf.String()
	case "RCFAST":
		return RCFAST, buf.String()
	case "RCSLOW":
		return RCSLOW, buf.String()
	case "XINPUT":
		return XINPUT, buf.String()
	case "XTAL1":
		return XTAL1, buf.String()
	case "XTAL2":
		return XTAL2, buf.String()
	case "XTAL3":
		return XTAL3, buf.String()
	case "PLL1X":
		return PLL1X, buf.String()
	case "PLL2X":
		return PLL2X, buf.String()
	case "PLL4X":
		return PLL4X, buf.String()
	case "PLL8X":
		return PLL8X, buf.String()
	case "PLL16X":
		return PLL16X, buf.String()

	// Variables
	case "RESULT":
		return RESULT, buf.String()

	// Flow Control
	case "ABORT":
		return ABORT, buf.String()
	case "CASE":
		return CASE, buf.String()
	case "IF":
		return IF, buf.String()
	case "IFNOT":
		return IFNOT, buf.String()
	case "NEXT":
		return NEXT, buf.String()
	case "QUIT":
		return QUIT, buf.String()
	case "REPEAT":
		return REPEAT, buf.String()
	case "RETURN":
		return RETURN, buf.String()

	// Memory
	case "BYTE":
		return BYTE, buf.String()
	case "WORD":
		return WORD, buf.String()
	case "LONG":
		return LONG, buf.String()
	case "BYTEFILL":
		return BYTEFILL, buf.String()
	case "WORDFILL":
		return WORDFILL, buf.String()
	case "LONGFILL":
		return LONGFILL, buf.String()
	case "BYTEMOVE":
		return BYTEMOVE, buf.String()
	case "WORDMOVE":
		return WORDMOVE, buf.String()
	case "LONGMOVE":
		return LONGMOVE, buf.String()
	case "LOOKUP":
		return LOOKUP, buf.String()
	case "LOOKUPZ":
		return LOOKUPZ, buf.String()
	case "LOOKDOWN":
		return LOOKDOWN, buf.String()
	case "LOOKDOWNZ":
		return LOOKDOWNZ, buf.String()
	case "STRSIZE":
		return STRSIZE, buf.String()
	case "STRCOMP":
		return STRCOMP, buf.String()

	// Directives
	case "STRING":
		return STRING, buf.String()
	case "CONSTANT":
		return CONSTANT, buf.String()
	case "FLOAT":
		return FLOAT, buf.String()
	case "ROUND":
		return ROUND, buf.String()
	case "TRUNC":
		return TRUNC, buf.String()
	case "FILE":
		return FILE, buf.String()

	// Locks
	case "LOCKNEW":
		return LOCKNEW, buf.String()
	case "LOCKRET":
		return LOCKRET, buf.String()
	case "LOCKCLR":
		return LOCKCLR, buf.String()
	case "LOCKSET":
		return LOCKSET, buf.String()

	// Chip Config
	case "CHIPVER":
		return CHIPVER, buf.String()
	case "CLKMODE":
		return CLKMODE, buf.String()
	case "_CLKMODE":
		return _CLKMODE, buf.String()
	case "CLKFREQ":
		return CLKFREQ, buf.String()
	case "_CLKFREQ":
		return _CLKFREQ, buf.String()
	case "CLKSET":
		return CLKSET, buf.String()
	case "_XINFREQ":
		return _XINFREQ, buf.String()
	case "_STACK":
		return _STACK, buf.String()
	case "_FREE":
		return _FREE, buf.String()

	// Registers
	case "CNT":
		return CNT, buf.String()
	case "CTRA":
		return CTRA, buf.String()
	case "CTRB":
		return CTRB, buf.String()
	case "DIRA":
		return DIRA, buf.String()
	case "DIRB":
		return DIRB, buf.String()
	case "INA":
		return INA, buf.String()
	case "INB":
		return INB, buf.String()
	case "OUTA":
		return OUTA, buf.String()
	case "OUTB":
		return OUTB, buf.String()
	case "FRQA":
		return FRQA, buf.String()
	case "FRQB":
		return FRQB, buf.String()
	case "PHSA":
		return PHSA, buf.String()
	case "PHSB":
		return PHSB, buf.String()
	case "VCFG":
		return VCFG, buf.String()
	case "VSCL":
		return VSCL, buf.String()
	case "PAR":
		return PAR, buf.String()
	case "SPR":
		return SPR, buf.String()

	// Process Control
	case "WAITCNT":
		return WAITCNT, buf.String()
	case "WAITPEQ":
		return WAITPEQ, buf.String()
	case "WAITPNE":
		return WAITPNE, buf.String()
	case "WAITVID":
		return WAITVID, buf.String()

	// Cog Control
	case "COGID":
		return COGID, buf.String()
	case "COGNEW":
		return COGNEW, buf.String()
	case "COGINIT":
		return COGINIT, buf.String()
	case "COGSTOP":
		return COGSTOP, buf.String()
	case "REBOOT":
		return REBOOT, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENTIFIER, buf.String()
}
