package parser

import (
	"github.com/unsafecast/quack/ast"
	"github.com/unsafecast/quack/errors"
	"github.com/unsafecast/quack/lexer"
	"github.com/unsafecast/quack/token"
)

type Parser struct {
	lexer *lexer.Lexer
}

func New(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}

func (parser *Parser) Ident() (ast.Node, error) {
	tok := parser.lexer.CurrentToken

	if tok.Kind != token.KIND_IDENT {
		return nil, parser.newSyntaxError("expected an identifier")
	}

	parser.advance()
	return ast.NewIdent(string(tok.Value), tok.Offset), nil
}

func (parser *Parser) Named() (ast.Node, error) {
	return parser.Ident() // TEMP
}

func (parser *Parser) NumLit() (ast.Node, error) {
	tok := parser.lexer.CurrentToken

	if tok.Kind != token.KIND_NUM_LIT {
		return nil, parser.newSyntaxError("expected a number literal")
	}

	parser.advance()
	return ast.NewNumLit(tok.Number(), tok.Offset), nil
}

func (parser *Parser) Expr() (ast.Node, error) {
	switch parser.lexer.CurrentToken.Kind {
	case token.KIND_IDENT:
		return parser.Named()
	case token.KIND_NUM_LIT:
		return parser.NumLit()
	default:
		return nil, parser.newSyntaxError("expected an expression")
	}
}

func (parser *Parser) Assignment() (ast.Node, error) {
	offset := parser.lexer.CurrentToken.Offset
	if parser.advance().Kind != token.KIND_LET {
		return nil, parser.newSyntaxError("expected 'let'")
	}

	name, err := parser.Named()
	if err != nil {
		return nil, err
	}

	if parser.advance().Kind != token.KIND_EQUALS {
		return nil, parser.newSyntaxError("expected '='")
	}

	val, err := parser.Expr()
	if err != nil {
		return nil, err
	}

	if parser.advance().Kind != token.KIND_SEMICOLON {
		return nil, parser.newSyntaxError("expected ';'")
	}

	return ast.NewAssignment(name, val, offset), nil
}

func (parser *Parser) Stmt() (ast.Node, error) {
	switch parser.lexer.CurrentToken.Kind {
	case token.KIND_LET:
		return parser.Assignment()
	default:
		return nil, parser.newSyntaxError("expected statement")
	}
}

func (parser *Parser) advance() token.Token {
	old := parser.lexer.CurrentToken
	parser.lexer.Next()
	return old
}

func (parser *Parser) newSyntaxError(msg string) error {
	return errors.NewSyntaxError(msg, parser.lexer.Offset)
}
