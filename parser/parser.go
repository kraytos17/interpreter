package parser

import (
	"fmt"

	"github.com/soumil-kumar17/interpreter/ast"
	"github.com/soumil-kumar17/interpreter/lexer"
	"github.com/soumil-kumar17/interpreter/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors    []string
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, errors: []string{}}
	parser.nextTok()
	parser.nextTok()
	return parser
}

func (parser *Parser) nextTok() {
	parser.currToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(t token.TokenType) {
	message := fmt.Sprintf("Unexpexted token %s, expected %s", t, parser.currToken.Type)
	parser.errors = append(parser.errors, message)
}

func (parser *Parser) ParseProg() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for parser.currToken.Type != token.EOF {
		stmt := parser.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		parser.nextTok()
	}
	return prog
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.currToken}

	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: parser.currToken, Value: parser.currToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	for !parser.currTokenIs(token.SEMICOLON) {
		parser.nextTok()
	}
	return stmt
}

func (parser *Parser) currTokenIs(t token.TokenType) bool {
	return parser.currToken.Type == t
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) expectPeek(t token.TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextTok()
		return true
	} else {
		parser.peekError(t)
		return false
	}
}
