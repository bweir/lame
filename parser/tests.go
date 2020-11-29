package parser

import "github.com/bweir/lame/token"

func isBlock(tok token.Token) bool {
	return tok.Type == token.PUB || tok.Type == token.PRI || tok.Type == token.CON || tok.Type == token.DAT || tok.Type == token.OBJ || tok.Type == token.VAR
}
