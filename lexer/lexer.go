package lexer

// TokenType represents the type of tokens.
type TokenType int

// Token types
const (
    EOF TokenType = iota
    IDENTIFIER
    DIGIT
    SENT
    CHAR
    DECI64
    FUNC
    MAIN
    DISPLAY
    DISPLAYLN
    DISPLAYF
    PLUS
    MINUS
    ASTERISK
    SLASH
    ASSIGN
    SEMICOLON
    LPAREN
    RPAREN
    LBRACE
    RBRACE
    COMMA
)

// Token holds the token type and its literal value.
type Token struct {
    Type    TokenType // The type of the token
    Literal string    // The literal value of the token
}

// Lexer holds the state of the lexer.
type Lexer struct {
    input        string // The input string
    position     int    // Current position in the input
    readPosition int    // Current reading position in the input
    ch           byte   // Current character being examined
}

// NewLexer initializes a new lexer instance.
func NewLexer(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar() // Read the first character
    return l
}

// readChar reads the next character and updates the lexer state.
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0 // 0 indicates EOF
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

// NextToken retrieves the next token from the input.
func (l *Lexer) NextToken() Token {
    var tok Token

    // Skip whitespace
    for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
        l.readChar()
    }

    switch l.ch {
    case 0:
        tok = Token{Type: EOF, Literal: ""}
    case '+':
        tok = newToken(PLUS)
    case '-':
        tok = newToken(MINUS)
    case '*':
        tok = newToken(ASTERISK)
    case '/':
        tok = newToken(SLASH)
    case '=':
        tok = newToken(ASSIGN)
    case ';':
        tok = newToken(SEMICOLON)
    case '(':
        tok = newToken(LPAREN)
    case ')':
        tok = newToken(RPAREN)
    case '{':
        tok = newToken(LBRACE)
    case '}':
        tok = newToken(RBRACE)
    case ',':
        tok = newToken(COMMA)
    default:
        if isLetter(l.ch) {
            literal := l.readIdentifier()
            tok.Literal = literal
            tok.Type = identifyKeyword(literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Literal = l.readNumber()
            tok.Type = DIGIT // You can modify this to differentiate decimal types
            return tok
        } else if l.ch == '"' {
            tok.Literal = l.readString() // Read a string
            tok.Type = SENT
            return tok
        }
    }

    l.readChar() // Read the next character
    return tok
}

// newToken creates a new token.
func newToken(tokenType TokenType) Token {
    return Token{Type: tokenType, Literal: string(tokenType)}
}

// identifyKeyword checks if a literal matches a keyword.
func identifyKeyword(literal string) TokenType {
    switch literal {
    case "func":
        return FUNC
    case "main":
        return MAIN
    case "display":
        return DISPLAY
    default:
        return IDENTIFIER
    }
}

// Utility functions to identify letters and digits
func isLetter(ch byte) bool {
    return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
    start := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
    start := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[start:l.position]
}

func (l *Lexer) readString() string {
    start := l.position + 1 // Move past the opening quote
    l.readChar()            // Read the first character after the quote
    for l.ch != '"' && l.ch != 0 {
        l.readChar()
    }
    return l.input[start:l.position]
}