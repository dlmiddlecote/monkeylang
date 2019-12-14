package ast

import "github.com/dlmiddlecote/monkeylang/token"

// Node defines an interface for all nodes in the AST.
type Node interface {
	TokenLiteral() string
}

// Statement defines the interface for all statement nodes.
type Statement interface {
	Node
	statementNode()
}

// Expression defines the interface for all expression nodes.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node. All programs consist of a slice of Statement(s)
type Program struct {
	Statements []Statement
}

// TokenLiteral implements Node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement represents a statement of the form
// let name = value
type LetStatement struct {
	Token token.Token // the let token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral implements Node
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement ...
type ReturnStatement struct {
	Token       token.Token // the return token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral implements Node
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// Identifier represents an identifer and holds the name of the identifier
type Identifier struct {
	Token token.Token // the ident token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral implements Node
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
