package ast

import (
	"strings"

	"kisumu/token"
)

// Node is an interface that all AST nodes implement.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is an interface for all statement nodes.
type Statement interface {
	Node
	statementNode()
}

// Expression is an interface for all expression nodes.
type Expression interface {
	Node
	expressionNode()
}

// Program represents the entire source code.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal representation of the first token.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns a string representation of the entire program.
func (p *Program) String() string {
	var out strings.Builder

	for _, stmt := range p.Statements {
		if stmt != nil {
			out.WriteString(stmt.String())
			out.WriteString("\n")
		}
	}

	return out.String()
}

// PackageStatement represents a package declaration.
type PackageStatement struct {
	Token       token.Token
	PackageName string
}

func (ps *PackageStatement) statementNode()       {}
func (ps *PackageStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PackageStatement) String() string {
	if ps == nil {
		return ""
	}
	return "PACK " + ps.PackageName
}

// ImportStatement represents an import statement.
type ImportStatement struct {
	Token      token.Token
	ModuleName string
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) String() string {
	if is == nil {
		return ""
	}
	return "IMPORT " + is.ModuleName
}

// LetStatement represents a variable declaration.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	if ls == nil {
		return ""
	}
	var out strings.Builder
	out.WriteString("LET ")
	if ls.Name != nil {
		out.WriteString(ls.Name.String())
	}
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	return out.String()
}

// Identifier represents a variable or function name.
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	if i == nil {
		return ""
	}
	return i.Value
}

// ExpressionStatement represents a statement consisting of an expression.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es == nil {
		return ""
	}
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// PackStatement represents a package statement.
type PackStatement struct {
	Token token.Token
	Name  string
}

func (ps *PackStatement) statementNode()       {}
func (ps *PackStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PackStatement) String() string {
	if ps == nil {
		return ""
	}
	return "PACK " + ps.Name
}

// CoreStatement represents a core block.
type CoreStatement struct {
	Token      token.Token
	Statements []Statement
}

func (cs *CoreStatement) statementNode()       {}
func (cs *CoreStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *CoreStatement) String() string {
	if cs == nil {
		return ""
	}
	var out strings.Builder
	out.WriteString("CORE {\n")

	for _, stmt := range cs.Statements {
		if stmt != nil {
			out.WriteString("\t")
			out.WriteString(stmt.String())
			out.WriteString("\n")
		}
	}

	out.WriteString("}")
	return out.String()
}

// DeclareStatement represents a declare statement.
type DeclareStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ds *DeclareStatement) statementNode()       {}
func (ds *DeclareStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DeclareStatement) String() string {
	if ds == nil {
		return ""
	}
	var out strings.Builder
	out.WriteString("DECLARE ")
	if ds.Name != nil {
		out.WriteString(ds.Name.String())
	}
	out.WriteString(" = ")
	if ds.Value != nil {
		out.WriteString(ds.Value.String())
	}
	return out.String()
}

// DisplayStatement represents a display statement.
type DisplayStatement struct {
	Token token.Token
	Name  *Identifier
}

func (ds *DisplayStatement) statementNode()       {}
func (ds *DisplayStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DisplayStatement) String() string {
	if ds == nil {
		return ""
	}
	if ds.Name != nil {
		return "DISPLAY " + ds.Name.String()
	}
	return ""
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
