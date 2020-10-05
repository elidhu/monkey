package lexer

import (
	"github.com/kevinglasson/monkey/token"
)

// NextToken examines generates the next Token and advances the indexer.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Skip the whitespace first!
	l.skipWhitespace()

	// Examine the current char.
	switch l.ch {
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '=':
		// If the next char is an '=' then we have and '=='
		if l.peekChar() == '=' {
			// Combine the current char and the next one.
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
			// Otherwise it's just an assignment
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		// If the next char is an '=' then we have and '!='
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NEQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	// We actually set this as the current char when we've reached the end of
	// the input string.
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// If we have a letter.
		if isLetter(l.ch) {
			// Read the identifier.
			tok.Literal = l.readIdentifier()
			// Determine it's type i.e. is it a keyword or just a user defined
			// identifier (variable name etc).
			tok.Type = token.LookupIdent(tok.Literal)
			// Return early as we have already advanced the char indexer.
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return tok
}

// skipWhitespace advances the char indexer until the whitespace is done.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// peekChar looks at the next char in the input without advancing the indexer.
func (l *Lexer) peekChar() byte {
	// If the next char will be the end of the input the return 0.
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	// Loop while we have letters.
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// isLetter determines if a character is a letter (or an underscore).
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	startPos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Lexer is the languages lexer struct.
type Lexer struct {
	// input is the input string to lex.
	input string
	// position is the current position in the input (current char).
	position int
	// readPosition is the reading position in the input (next char).
	readPosition int
	// ch is the current char under examination.
	ch byte
}

// New creates a new lexer for an input string.
func New(input string) *Lexer {
	// Init a lexer.
	l := &Lexer{input: input}
	// Read teh first char of the string.
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// Read the next char
		l.ch = l.input[l.readPosition]
	}
	// Advance the current, and next char indexes
	l.position = l.readPosition
	l.readPosition++
}
