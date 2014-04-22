package lexer

import (
	"unicode/utf8"
)

const Eof = -1

type TokenType int

type Token struct {
	Type TokenType
	Value string
}

type Lexer struct {
	input string
	start, pos, width int
	tokens chan Token
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		tokens: make(chan Token),
	}
}

type StateFn func(*Lexer) StateFn

func (lex *Lexer) Run(start StateFn) {
	for state := start; state != nil; {
		state = state(lex)
	}
	close(lex.tokens)
}

func (lex *Lexer) NextRune() rune {
	if lex.pos >= len(lex.input) {
		lex.width = 0
		return Eof
	}
	r, w := utf8.DecodeRuneInString(lex.input[lex.pos:])
	lex.width = w
	lex.pos += w
	return r
}

// Back up one rune. Can only be called once per call of NextRune().
func (lex *Lexer) Backup() {
	lex.pos -= lex.width
}

func (lex *Lexer) Emit(t TokenType) {
	lex.tokens <- Token{t, lex.input[lex.start:lex.pos]}
	lex.start = lex.pos
}

func (lex *Lexer) Next() (Token, bool) {
	t, ok := <-lex.tokens
	return t, ok
}
