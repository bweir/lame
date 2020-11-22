package lexer

import (
	"bytes"
	"container/list"
	"fmt"

	"github.com/bweir/lame/token"
)

func (s *Scanner) getSpace() (buf bytes.Buffer) {
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

	// fmt.Println("space", buf.String())

	// for i := 0; i < buf.Len(); i++ {
	// 	s.unread()
	// }

	// fmt.Printf("BUFFER '%s'", buf.String())

	return buf
}

func (s *Scanner) scanSpace() (tok token.Token) {

	buf := s.getSpace()
	// for i := 0; i < buf.Len(); i++ {
	// 	fmt.Println(buf.String())
	// 	s.read()
	// }
	return s.makeToken(token.NULL, buf.String())
}

func (s *Scanner) getCurrentIndent() (currentIndent int) {
	for e := s.indent.Front(); e != nil; e = e.Next() {
		currentIndent += e.Value.(int)
	}
	return currentIndent
}

func printList(li *list.List) {
	for e := li.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d, ", e.Value)
	}
	fmt.Printf("\n")
}

func (s *Scanner) scanIndentLevel() (tok token.Token) {
	ch := s.read()
	if ch == eof {
		return s.makeToken(token.EOF, "")
	}
	var buf bytes.Buffer
	if isSpace(ch) {
		s.unread()
		buf = s.getSpace()
	}

	fmt.Printf("BFF '%s'\n", buf.String())
	fmt.Printf("lien '%t'\n", s.blockStart)

	newIndent := buf.Len()

	if s.blockStart {
		s.indent.PushBack(newIndent)
		s.blockStart = false
		s.lineStart = false
		printList(s.indent)
		return s.makeToken(token.NULL, buf.String())
	}

	currentIndent := s.indent.Back().Value.(int)

	printList(s.indent)
	fmt.Println("current indent level", currentIndent)
	fmt.Println("    new indent level", newIndent)
	if newIndent > currentIndent {
		fmt.Println("indenting")
		printList(s.indent)
		s.indent.PushBack(newIndent - currentIndent)
		currentIndent = s.indent.Back().Value.(int)
		for i := 0; i < len(buf.String()); i++ {
			s.read()
			s.column++
		}
		printList(s.indent)
		s.lineStart = false
		return s.makeToken(token.INDENT, buf.String())
	}

	if newIndent < currentIndent {
		fmt.Println("dedenting")
		fmt.Println(newIndent - currentIndent)
		printList(s.indent)
		onestep := s.indent.Remove(s.indent.Back()).(int)
		for i := 0; i < onestep; i++ {
			s.read()
		}
		printList(s.indent)
		return s.makeToken(token.DEDENT, buf.String())
	}

	s.lineStart = false
	return s.makeToken(token.NULL, buf.String())
}
