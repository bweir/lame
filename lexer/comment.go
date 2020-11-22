package lexer

import (
	"bytes"

	"github.com/bweir/lame/token"
)

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
