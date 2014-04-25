package lexer

import "testing"

func AssertStream(t *testing.T, tokenize func(string) *Lexer, src string, produceStream func(chan Token)) {
	stream := make(chan Token)
	go func() {
		produceStream(stream)
		close(stream)
	}()

	expecteds := []Token{}
	for expected := range stream {
		expecteds = append(expecteds, expected)
	}

	tokens := tokenize(src)
	actuals := []Token{}
	for {
		actual, ok := tokens.Next()
		if !ok {
			break
		}
		actuals = append(actuals, actual)
	}

	for i, expected := range expecteds {
		if actuals[i] != expected {
			t.Fatalf(`Unexpected tokenization:
Expected: %s
Actual: %s
Source:
-------
%s`, expecteds, actuals, src)
			break
		}
	}
}
