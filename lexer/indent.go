package lexer

import (
	"bytes"
	"container/list"
	"fmt"

	"github.com/bweir/lame/token"
)

func printList(li *list.List) {
	for e := li.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d, ", e.Value)
	}
	fmt.Printf("\n")
}

func (s *Scanner) readIndent() {
	var buf bytes.Buffer

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

	s.newIndent = buf.Len()
	if s.blockStart {
		s.blockStart = false
	}
}

func (s *Scanner) scanIndent() (tok token.Token) {
	s.indent.PushBack(s.newIndent)
	return s.makeToken(token.INDENT, "")
}

func (s *Scanner) scanDedent() (tok token.Token) {
	s.indent.Remove(s.indent.Back())
	return s.makeToken(token.DEDENT, "")
}
