package ast

import (
	"bytes"

	"github.com/kevinglasson/monkey/token"
)

// Node is the node of the AST.
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

// Expression is something that produces a value.
type Expression interface {
	Node
	expressionNode()
}

// Program is a list of statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	// Loop through all statements in the program and write them to the output
	// buffer.
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Identifier is a struct representing the IDENT token with it's value.
type Identifier struct {
	Token token.Token
	Value string
}

// expressionNode implements Expression for Identifier.
func (i *Identifier) expressionNode() {}

// TokenLiteral implements Node for Identifier.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// LetStatement implements Node for the LET statement.
type LetStatement struct {
	// The LET token.
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral implements part of the Node interface so we can output this
// statement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) statementNode()       {}

// String implements part of the Node interface so we can output this statement
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement implements Node for the RETURN statement.
type ReturnStatement struct {
	// The RETURN token.
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral implements Node for ReturnStatement.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) statementNode()       {}

// String implements part of the Node interface so we can output this statement
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement implements Node, the purpose of this is so that we
// fulfill the Node interface and add expressions to the AST.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// TokenLiteral implements Node for ExpressionStatement.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) statementNode()       {}

// String implements part of the Node interface so we can output this statement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
