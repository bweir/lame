package lexer

import (
	"bytes"
	"strings"

	"github.com/bweir/lame/token"
	"github.com/bweir/lame/token/state"
)

func (s *Scanner) scanIdentifier() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isIdentifier(ch) && !isDecimalDigit(ch) {
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
		s.state = state.CONSTANT
		s.blockStart = true
		return s.makeToken(token.CON, buf.String())
	case "DAT":
		s.state = state.DATA
		s.blockStart = true
		return s.makeToken(token.DAT, buf.String())
	case "OBJ":
		s.state = state.OBJECT
		s.blockStart = true
		return s.makeToken(token.OBJ, buf.String())
	case "PRI":
		s.state = state.FUNCTION
		s.blockStart = true
		return s.makeToken(token.PRI, buf.String())
	case "PUB":
		s.state = state.FUNCTION
		s.blockStart = true
		return s.makeToken(token.PUB, buf.String())
	case "VAR":
		s.state = state.VARIABLE
		s.blockStart = true
		return s.makeToken(token.VAR, buf.String())

	// Constants
	case "TRUE":
		return s.makeToken(token.TRUE, buf.String())
	case "FALSE":
		return s.makeToken(token.FALSE, buf.String())

	// Flow Control
	case "CASE":
		return s.makeToken(token.CASE, buf.String())
	case "IF":
		return s.makeToken(token.IF, buf.String())
	case "ELSEIF":
		return s.makeToken(token.ELSEIF, buf.String())
	case "ELSE":
		return s.makeToken(token.ELSE, buf.String())
	case "NEXT":
		return s.makeToken(token.NEXT, buf.String())
	case "QUIT":
		return s.makeToken(token.QUIT, buf.String())
	case "REPEAT":
		return s.makeToken(token.REPEAT, buf.String())
	case "FROM":
		return s.makeToken(token.FROM, buf.String())
	case "TO":
		return s.makeToken(token.TO, buf.String())
	case "STEP":
		return s.makeToken(token.STEP, buf.String())
	case "WHILE":
		return s.makeToken(token.WHILE, buf.String())
	case "UNTIL":
		return s.makeToken(token.UNTIL, buf.String())
	case "RETURN":
		return s.makeToken(token.RETURN, buf.String())

	// Memory
	case "BYTE":
		return s.makeToken(token.BYTE, buf.String())
	case "WORD":
		return s.makeToken(token.WORD, buf.String())
	case "LONG":
		return s.makeToken(token.LONG, buf.String())

	// logical
	case "NOT":
		return s.makeToken(token.NOT, buf.String())
	case "AND":
		return s.makeToken(token.AND, buf.String())
	case "OR":
		return s.makeToken(token.OR, buf.String())
	}

	return s.makeToken(token.IDENTIFIER, buf.String())
}
