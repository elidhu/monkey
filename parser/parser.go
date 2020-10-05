package parser

import (
	"fmt"

	"github.com/kevinglasson/monkey/ast"
	"github.com/kevinglasson/monkey/lexer"
	"github.com/kevinglasson/monkey/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS  // == LESSGREATER // > or <
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL    // myFunction(X)
)

// Parser is a parser for the programming language.
type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// New creates a new initialised Parser from a Lexer.
func New(l *lexer.Lexer) *Parser {
	// Create a new parser.
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Register parse functions
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)

	// Initialise curToken and peekToken by reading two tokens.
	p.nextToken()
	p.nextToken()

	return p
}

// Errors returns all of the errors the Parser has collected.
func (p *Parser) Errors() []string {
	return p.errors
}

// ParseProgram parses the program.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	program.Statements = []ast.Statement{}

	// Until the current token is EOF.
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// peekError generates a peek error and adds it to the parsers errors slice.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken indexes the parser to the next token.
func (p *Parser) nextToken() {
	// The current token is not the peek token.
	p.curToken = p.peekToken
	// The peek token is the next token generated from the lexer.
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// We expect an IDENTIFIER immediately following a LET.
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Set the Name (Identifier) for the Statement.
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// We expect an ASSIGN after the LET IDENTIFIER sequence.
	if !p.expectPeek(token.ASSIGN) {
		p.nextToken()
	}

	// TODO: Skipping expressions until we get to a SEMICOLON (just for now).
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: Skipping expressions until we get to a SEMICOLON (just for now).
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek advances the token indexer if the TokenType is as expected.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	// Add an error and return.
	p.peekError(t)
	return false
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
