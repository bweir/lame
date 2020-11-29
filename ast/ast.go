package ast

import "github.com/bweir/lame/token"

// base interface of the AST
type Node interface {
	Pos() token.Pos
}

// Blocks
type Block interface {
	Node
	blockNode()
}

// Statements
type Statement interface {
	Node
	statementNode()
}

// Declaration
type Declaration interface {
	Node
	declarationNode()
}

// Block definitions

type (
	ConBlock struct {
		From, To     token.Pos
		Declarations []ConstantDeclaration
	}

	DatBlock struct{ From, To token.Pos }
	ObjBlock struct{ From, To token.Pos }
	PriBlock struct{ From, To token.Pos }
	PubBlock struct{ From, To token.Pos }
	VarBlock struct{ From, To token.Pos }
)

func (b *ConBlock) Pos() token.Pos { return b.From }
func (b *DatBlock) Pos() token.Pos { return b.From }
func (b *ObjBlock) Pos() token.Pos { return b.From }
func (b *PriBlock) Pos() token.Pos { return b.From }
func (b *PubBlock) Pos() token.Pos { return b.From }
func (b *VarBlock) Pos() token.Pos { return b.From }

func (*ConBlock) blockNode() {}
func (*DatBlock) blockNode() {}
func (*ObjBlock) blockNode() {}
func (*PriBlock) blockNode() {}
func (*PubBlock) blockNode() {}
func (*VarBlock) blockNode() {}

// Statement definitions

type (
	ConStatement struct{ From, To token.Pos }
)

func (s *ConStatement) Pos() token.Pos { return s.From }

func (*ConStatement) statementNode() {}

// Statement definitions

type (
	ConstantDeclaration struct {
		From, To token.Pos
		Name     string
		Value    string
	}
)

func (s *ConstantDeclaration) Pos() token.Pos { return s.From }

func (*ConstantDeclaration) declarationNode() {}

// An object represents a Lame object.
type Object struct {
	Blocks []Block
}
