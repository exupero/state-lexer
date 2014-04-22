package lexer

import "testing"

func AssertStream(t *testing.T, src string, tokenize func(string) *Lexer, produceStream func(chan Token)) {
	stream := make(chan Token)
	go func() {
		produceStream(stream)
		close(stream)
	}()

	tokens := tokenize(src)
	i := 1
	for expected := range stream {
		actual, ok := tokens.Next()
		if !ok {
			t.Fatalf("tokenization of `%s` ended prematurely", src)
		}
		if actual != expected {
			t.Errorf("tokenization of `%s` does not produce %s but rather %s (token %d)", src, expected, actual, i)
		}
		i++
	}
	if token, ok := tokens.Next(); ok {
		t.Fatalf("tokenization of `%s` produced unexpected token %s", src, token)
	}
}
