package lexer

import "github.com/bweir/lame/token"

func (s *Scanner) scanOperator() (tok token.Token) {
	ch := s.read()

	switch ch {
	case eof:
		return s.makeToken(token.EOF, "")
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
