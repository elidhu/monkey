package parser

import (
	"fmt"
	"testing"

	"github.com/kevinglasson/monkey/ast"
	"github.com/kevinglasson/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 260991;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got %d",
			len(program.Statements),
		)
	}

	// Create an slice of expected identifiers.
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	// For each expected identifier check that it is as expected in the
	// programs statement slice.
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	// If the parser did have errors then log them and fail!
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser errors: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// Check the token's literal value is "let"
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	// Assert the Type of this Statement is a LetStatement
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	fmt.Printf("%+v", letStmt)

	// Test that the Value of the Name (Identifier) of this Statement is as
	// expected.
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	// Test that the
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf(
			"letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name,
			letStmt.Name.TokenLiteral(),
		)
		return false
	}
	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	// Assert type of each statement is ReturnStatement.
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got %T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf(
				"returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral(),
			)
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program has not engouh statements. got=%d",
			len(program.Statements),
		)
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. gto %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}
