package lexer

import (
	"bytes"
	"strconv"

	"github.com/bweir/lame/token"
)

func (s *Scanner) scanString() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == '"' {
			break
		} else if ch == eof {
			return s.makeToken(token.UNEXPECTED_EOF, "")
		} else if ch == '\\' {
			if ch := s.read(); ch == eof {
				break
			} else if ch == 'n' {
				_, _ = buf.WriteRune('\n')
			} else if ch == 't' {
				_, _ = buf.WriteRune('\t')
			} else if ch == '"' {
				_, _ = buf.WriteRune('"')
			} else if isDigit(ch) {
				s.unread()
				tok := s.scanDecimalNumber()
				num, _ := strconv.Atoi(tok.Literal)
				buf.WriteRune(rune(num))
			} else {
				return s.makeToken(token.ILLEGAL, string(ch))
			}
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.STRING, buf.String())
}
