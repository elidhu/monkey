package ast

import "github.com/kevinglasson/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement implements Node for the let statement.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// statementNode implements Statement for LetStatement.
func (ls *LetStatement) statementNode() {}

// TokenLiteral implements Node for LetStatement.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is a struct representing the IDENT token with it's value.
type Identifier struct {
	Token token.Token
	Value string
}

// expressionNode implements Expression for Identifier.
func (i *Identifier) expressionNode() {}

// TokenLiteral implements Node for Identifier.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
