package lexer

import (
	"bytes"

	"github.com/bweir/lame/token"
)

func (s *Scanner) scanBinaryNumber() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isBinaryDigit(ch) && !isGroupSeparator(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.BINARY_NUMBER, buf.String())
}

func (s *Scanner) scanQuaternaryNumber() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isQuaternaryDigit(ch) && !isGroupSeparator(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.QUATERNARY_NUMBER, buf.String())
}

func (s *Scanner) scanDecimalNumber() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDecimalDigit(ch) && !isGroupSeparator(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.DECIMAL_NUMBER, buf.String())
}

func (s *Scanner) scanHexadecimalNumber() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isHexadecimalDigit(ch) && !isGroupSeparator(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return s.makeToken(token.HEXADECIMAL_NUMBER, buf.String())
}
