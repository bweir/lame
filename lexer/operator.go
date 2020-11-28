package lexer

import "github.com/bweir/lame/token"

func (s *Scanner) scanOperator() (tok token.Token) {
	ch := s.read()

	switch ch {
	case eof:
		return s.makeToken(token.EOF, "")
	case ',':
		return s.makeToken(token.COMMA, string(ch))

	case ':':
		return s.makeToken(token.COLON, string(ch))
	case '.':
		return s.makeToken(token.DOT, string(ch))
	case '|':
		return s.makeToken(token.PIPE, string(ch))
	case '@':
		return s.makeToken(token.AT, string(ch))
	case '#':
		return s.makeToken(token.POUND, string(ch))

	// Bitwise
	case '!':
		return s.makeToken(token.BITWISE_NOT, string(ch))

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
