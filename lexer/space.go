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

func (s *Scanner) scanIndentLevel() (tok token.Token) {
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

	// fmt.Printf("BFF '%s'\n", buf.String())
	// fmt.Printf("lien '%t'\n", s.blockStart)

	newIndent := buf.Len()

	if s.blockStart {
		s.indent.PushBack(newIndent)
		s.blockStart = false
		s.lineStart = false
		// printList(s.indent)
		return s.makeToken(token.NULL, buf.String())
	}

	currentIndent := s.indent.Back().Value.(int)

	// printList(s.indent)
	// fmt.Println("current indent level", currentIndent)
	// fmt.Println("    new indent level", newIndent)
	if newIndent > currentIndent {
		// fmt.Println("indenting")
		// printList(s.indent)
		s.indent.PushBack(newIndent)
		// printList(s.indent)
		s.lineStart = false
		return s.makeToken(token.INDENT, buf.String())
	}

	if newIndent < currentIndent {
		// fmt.Println("dedenting")
		// printList(s.indent)
		lastStep := s.indent.Remove(s.indent.Back()).(int)
		for i := 0; i < lastStep-s.indent.Back().Value.(int); i++ {
			s.unread()
		}
		// printList(s.indent)
		return s.makeToken(token.DEDENT, buf.String())
	}

	s.lineStart = false
	return s.makeToken(token.NULL, buf.String())
}
