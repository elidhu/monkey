package token

// TokenType is a constant representing the token types that are lexed.
type TokenType string

// Token is a struct to package a lexed token type with it's literal value.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Other

	// ILLEGAL is the TokenType for unknown tokens.
	ILLEGAL = "ILLEGAL"
	// EOF is the TokenType to mark the end of a file.
	EOF = "EOF"

	// Identifiers and literals.

	// IDENT is the TokenType to mark an identifier.
	IDENT = "IDENT"
	// INT is the TokenType for integer numbers.
	INT = "INT"
	// TRUE
	TRUE = "TRUE"
	// FALSE
	FALSE = "FALSE"
	// IF
	IF = "IF"
	// ELSE
	ELSE = "ELSE"
	// RETURN
	RETURN = "RETURN"

	// Operators.

	// ASSIGN is the TokenType for an assignment.
	ASSIGN = "="
	// PLUS is the TokenType for the addition operation.
	PLUS = "+"
	// MINUS is the TokenType for the subtraction operation.
	MINUS = "-"
	// BANG is the TokenType for the not operator.
	BANG = "!"
	// ASTERISK is the TokenType for the multiplication operation.
	ASTERISK = "*"
	// SLASH is the TokenType for the  division operation.
	SLASH = "/"
	// LT is the TokenType for the less than comparison.
	LT = "<"
	// GT is the TokenType for the greater than comparison.
	GT = ">"
	// EQ is the TokenType for the equality operator.
	EQ = "=="
	// NEQ is the TokenType for the not equal operator.
	NEQ = "!="

	// Delimiters.

	// COMMA is the TokenType to delimit things.
	COMMA = ","
	// SEMICOLON is the TokenType to mark the termination of a statement.
	SEMICOLON = ";"

	// LPAREN is the TokenType to mark the left parentheses.
	LPAREN = "("
	// RPAREN is the TokenType to mark the right parentheses.
	RPAREN = ")"
	// LBRACE is the TokenType to mark the left brace.
	LBRACE = "{"
	// RBRACE is the TokenType to mark the right brace.
	RBRACE = "}"

	// Keywords.

	// FUNCTION is the TokenType to mark the function (fn) keyword.
	FUNCTION = "FUNCTION"
	// LET is the TokenType to mark the let keyword.
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent looks up the identifier and returns it's token type.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
