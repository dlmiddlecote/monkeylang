package ast

import (
	"bytes"

	"github.com/dlmiddlecote/monkeylang/token"
)

// Node defines an interface for all nodes in the AST.
type Node interface {
	TokenLiteral() string
	String() string
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

// String implements Node
func (p *Program) String() string {
	var out bytes.Buffer

	for i := range p.Statements {
		out.WriteString(p.Statements[i].String())
	}

	return out.String()
}

// LetStatement represents a statement of the form
// let name = value;
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

// String implements Node
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement represents a statement of the form
// return ReturnValue;
type ReturnStatement struct {
	Token       token.Token // the return token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral implements Node
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String implements Node
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())
	out.WriteString(" ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is a statement made of just one expression, basically a wrapper
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral implements Node
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String implements Node
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
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

// String implements Node
func (i *Identifier) String() string {
	return i.Value
}
