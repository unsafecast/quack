package lexer

import (
	"unicode"

	"github.com/unsafecast/quack/token"
)

type Lexer struct {
	CurrentToken token.Token
	PeekToken    token.Token

	Offset     int64
	usedOffset int64

	source   []rune
	currChar rune
}

func New(source []rune) *Lexer {
	lexer := &Lexer{
		source:     source,
		currChar:   source[0],
		Offset:     0,
		usedOffset: 0,
	}

	// - This is so we have a valid state
	// - After these calls, CurrentToken and PeekToken
	//   will be properly initialized
	lexer.Next()
	lexer.Next()

	return lexer
}

func (lexer *Lexer) Next() token.Token {
	for unicode.IsSpace(lexer.currChar) {
		lexer.advance()
	}

	lexer.usedOffset = lexer.Offset
	the := token.Invalid(lexer.currChar, lexer.usedOffset)

	switch lexer.currChar {
	case '=':
		lexer.advance()
		the = token.FromKind(token.KIND_EQUALS, lexer.usedOffset)

	case ';':
		lexer.advance()
		the = token.FromKind(token.KIND_SEMICOLON, lexer.usedOffset)

	default:
		if unicode.IsLetter(lexer.currChar) {
			the = lexer.Ident()
		} else if unicode.IsNumber(lexer.currChar) {
			the = lexer.Number()
		}
	}

	lexer.CurrentToken = lexer.PeekToken
	lexer.PeekToken = the

	return lexer.CurrentToken
}

func (lexer *Lexer) Ident() token.Token {
	start := lexer.Offset
	for isIdent(lexer.currChar) {
		lexer.advance()
	}

	val := lexer.source[start:lexer.Offset]
	switch string(val) {
	case "let":
		return token.FromKind(token.KIND_LET, lexer.usedOffset)
	default:
		return token.Ident(val, lexer.usedOffset)
	}
}

func (lexer *Lexer) Number() token.Token {
	start := lexer.Offset
	for isPartOfNumber(lexer.currChar) {
		lexer.advance()
	}
	return token.NumLit(lexer.source[start:lexer.Offset], lexer.usedOffset)
}

func (lexer *Lexer) advance() {
	lexer.Offset += 1
	if int(lexer.Offset) < len(lexer.source) {
		lexer.currChar = lexer.source[lexer.Offset]
	} else {
		lexer.currChar = '\000'
	}
}

func isIdent(what rune) bool {
	return unicode.IsLetter(what) ||
		unicode.IsNumber(what) ||
		what == '_'
}

func isPartOfNumber(what rune) bool {
	return unicode.IsNumber(what) || what == '.'
}
