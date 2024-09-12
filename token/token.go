package token

// TokenType represents the type of a token
type TokenType string

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Define token types as constants
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	PACKAGE = "PACKAGE"
	IMPORT  = "IMPORT"

	// Identifiers and literals
	IDENT  = "IDENT"
	DIGIT  = "DIGIT"
	STRING = "STRING"

	// Operators
	ADD          = "+"
	SUBTRACT     = "-"
	MULTIPLY     = "*"
	DIVIDE       = "/"
	LESS_THAN    = "<"
	GREATER_THAN = ">"
	ASSIGN       = "="

	// Delimiters
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	COMMA     = ","
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	CORE         = "CORE"
	IF_CASE      = "IF CASE"
	OTHERWISE    = "OTHERWISE"
	LOOP         = "LOOP"
	OPTION       = "OPTION"
	OPTION_CASE  = "OPTION CASE"
	DEFAULT_CASE = "DEFAULT CASE"
	RETURN       = "RETURN"
	EXIT         = "EXIT"
	NEXT         = "NEXT"
	METHOD       = "METHOD"
	DISPLAY      = "DISPLAY"
	DISPLAYLN    = "DISPLAYLN"
	DISPLAYF     = "DISPLAYF"
	DECLARE      = "DECLARE"
	CONSTANT     = "CONSTANT"
	DEFINE       = "DEFINE"
	PACK         = "PACK"
	BRING        = "BRING"
	DELAY        = "DELAY"
	CHOOSE       = "CHOOSE"
	CHANNEL      = "CHANNEL"
	FACE         = "FACE"
	STRUCTURE    = "STRUCTURE"
	MAPPING      = "MAPPING"
	ITERATE      = "ITERATE"
)

// NewToken creates a new token of the given type and literal value
func NewToken(tokenType TokenType, ch byte, line int, column int) Token {
	return Token{Type: tokenType, Literal: string(ch), Line: line, Column: column}
}
