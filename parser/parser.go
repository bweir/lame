package parser

import (
	"fmt"
	"io"

	"github.com/bweir/lame/ast"
	"github.com/bweir/lame/lexer"
	"github.com/bweir/lame/token"
)

type Parser struct {
	s   *lexer.Scanner
	buf struct {
		tok token.Token // last read token
		n   int         // buffer size (max=1)
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: lexer.NewScanner(r)}
}

func (p *Parser) scan() (tok token.Token) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok
	}

	tok = p.s.Scan()

	p.buf.tok = tok
	return
}

func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) scanIgnoreWhitespace() (tok token.Token) {
	tok = p.scan()
	for tok.Type == token.SPACE || tok.Type == token.NEWLINE {
		tok = p.scan()
	}
	return
}

func (p *Parser) Parse() (*ast.Object, error) {
	object := &ast.Object{}

	var list []ast.Block
	block, err := p.parseBlock()
	if err != nil {
		fmt.Println(err)
	}
	list = append(list, block)
	object.Blocks = list
	return object, nil
}

func (p *Parser) parseBlock() (block ast.Block, err error) {
	tok := p.scanIgnoreWhitespace()
	fmt.Printf("%q\n", tok)
	if !isBlock(tok) {
		return nil, fmt.Errorf("expecting block, I guess, found %q", tok.Literal)
	}

	switch tok.Type {
	case token.CON:
		p.unscan()
		block, err = p.parseConBlock()
	case token.OBJ:
		block = &ast.ObjBlock{}
	case token.PUB:
		block = &ast.PubBlock{}
	}
	return block, nil
}

func (p *Parser) parseConBlock() (block *ast.ConBlock, err error) {
	tok := p.scanIgnoreWhitespace()
	fmt.Printf("%q\n", tok)

	block = &ast.ConBlock{}
	for {
		tok = p.scanIgnoreWhitespace()
		if tok.Type != token.IDENTIFIER {
			return nil, fmt.Errorf("found %q, expected identifier", tok.Literal)
		}
		name := tok.Literal

		tok = p.scanIgnoreWhitespace()
		if tok.Type != token.ASSIGN {
			return nil, fmt.Errorf("found %q, expected assignment", tok.Literal)
		}

		tok = p.scanIgnoreWhitespace()
		if tok.Type != token.DECIMAL_NUMBER {
			return nil, fmt.Errorf("found %q, expected assignment", tok.Literal)
		}
		value := tok.Literal
		decl := &ast.ConstantDeclaration{Name: name, Value: value}
		block.Declarations = append(block.Declarations, *decl)

		tok = p.scanIgnoreWhitespace()
		p.unscan()
		if tok.Type != token.IDENTIFIER {
			break
		}
	}

	return
}
