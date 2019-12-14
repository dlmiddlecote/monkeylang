package token

const (
	// ILLEGAL denotes an illegal token
	ILLEGAL = "ILLEGAL"
	// EOF denotes the end of the input
	EOF = "EOF"

	//
	// Identifiers + literals
	//

	// IDENT is an identifier e.g. foo, add, x
	IDENT = "IDENT"
	// INT is an integer
	INT = "INT"

	//
	// Operators
	//

	// ASSIGN is the assignment operator
	ASSIGN = "="
	// PLUS is the addition operator
	PLUS = "+"
	// MINUS is the subtraction operator
	MINUS = "-"
	// ASTERISK is the multiplication operator
	ASTERISK = "*"
	// SLASH is the division operator
	SLASH = "/"
	// BANG is the not operator
	BANG = "!"

	//
	// Comparison operators
	//

	// EQ is the equality operator
	EQ = "=="
	// NOTEQ is the not equals operator
	NOTEQ = "!="
	// LT is the less than operator
	LT = "<"
	// GT is the greater than operator
	GT = ">"

	//
	// Delimiters
	//

	// COMMA is a comma
	COMMA = ","
	// SEMICOLON is a semicolon
	SEMICOLON = ";"
	// LPAREN is a left parenthesis
	LPAREN = "("
	// RPAREN is a right parenthesis
	RPAREN = ")"
	// LBRACE is a left brace
	LBRACE = "{"
	// RBRACE is a right brace
	RBRACE = "}"

	//
	// Keywords
	//

	// FUNCTION is the `fn` keyword (function)
	FUNCTION = "FUNCTION"
	// LET is the `let` keyword
	LET = "LET"
	// TRUE is the `true` keyword
	TRUE = "TRUE"
	// FALSE is the `false` keyword
	FALSE = "FALSE"
	// IF is the `if` keyword
	IF = "IF"
	// ELSE is the `else` keyword
	ELSE = "ELSE"
	// RETURN is the `return` keyword
	RETURN = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns the Type of the given identifier
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Type defines a type of token.
type Type string

// Token defines a token in our language, it's type and it's value.
type Token struct {
	Type    Type
	Literal string
}
