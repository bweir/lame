package lexer

import (
	"bytes"

	"github.com/bweir/lame/token"
)

func (s *Scanner) scanDecimalNumber() (tok token.Token) {
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

	return s.makeToken(token.DECIMAL_NUMBER, buf.String())
}
