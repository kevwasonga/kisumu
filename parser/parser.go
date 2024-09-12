package parser

import (
	"fmt"

	"kisumu/ast"
	"kisumu/lexer"
	"kisumu/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead at line %d, column %d",
		t, p.peekToken.Type, p.peekToken.Line, p.peekToken.Column)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
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
	case token.PACK:
		return p.parsePackStatement()
	case token.CORE:
		return p.parseCoreStatement()
	case token.DECLARE:
		return p.parseDeclareStatement()
	case token.DISPLAY:
		return p.parseDisplayStatement()
	default:
		return nil
	}
}

func (p *Parser) parsePackStatement() *ast.PackageStatement {
	stmt := &ast.PackageStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.PackageName = p.curToken.Literal

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	for p.curToken.Type != token.RBRACE && p.curToken.Type != token.EOF {
		p.nextToken()
		if p.curToken.Type == token.CORE || p.curToken.Type == token.DECLARE || p.curToken.Type == token.DISPLAY {
			stmtBody := p.parseStatement()
			if stmtBody != nil {
				// Handle the parsed statement within the pack body
				// You might need to update the ast.PackageStatement to include a slice of statements
			}
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return stmt
}

func (p *Parser) parseCoreStatement() *ast.CoreStatement {
	stmt := &ast.CoreStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	for p.curToken.Type != token.RBRACE && p.curToken.Type != token.EOF {
		p.nextToken()
		if p.curToken.Type == token.DECLARE || p.curToken.Type == token.DISPLAY {
			stmtBody := p.parseStatement()
			if stmtBody != nil {
				// Handle the parsed statement within the core body
				// You might need to update the ast.CoreStatement to include a slice of statements
			}
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return stmt
}

func (p *Parser) parseDeclareStatement() *ast.DeclareStatement {
	stmt := &ast.DeclareStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression()

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseDisplayStatement() *ast.DisplayStatement {
	stmt := &ast.DisplayStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseExpression() ast.Expression {
	// Implement expression parsing
	return nil
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
