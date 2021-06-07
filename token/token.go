package token

import (
	"fmt"
	"strconv"
)

type Kind int

const (
	KIND_INVALID Kind = iota
	KIND_EOF
	KIND_IDENT
	KIND_NUM_LIT
	KIND_EQUALS
	KIND_SEMICOLON
	KIND_LET
)

type Token struct {
	Kind   Kind
	Value  []rune
	Offset int64
}

//go:generate stringer -type=Kind

func (tok *Token) Number() float64 {
	if tok.Kind != KIND_NUM_LIT {
		panic(fmt.Errorf("token.Number() called on a non-number token"))
	}

	the, err := strconv.ParseFloat(string(tok.Value), 64)
	if err != nil {
		panic(fmt.Errorf("number token has non-number value"))
	}

	return the
}

func (tok *Token) String() string {
	return tok.Kind.String() + "\t: '" + string(tok.Value) + "'"
}

func FromKind(kind Kind, offset int64) Token {
	return Token{
		Kind:   kind,
		Offset: offset,
	}
}

func Invalid(val rune, offset int64) Token {
	return Token{
		Kind:   KIND_INVALID,
		Value:  []rune{val},
		Offset: offset,
	}
}

func Ident(val []rune, offset int64) Token {
	return Token{
		Kind:   KIND_IDENT,
		Value:  val,
		Offset: offset,
	}
}

func NumLit(val []rune, offset int64) Token {
	return Token{
		Kind:   KIND_NUM_LIT,
		Value:  val,
		Offset: offset,
	}
}
