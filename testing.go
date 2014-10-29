package lexer

import "testing"

func AssertStream(t *testing.T, tokenize func(string) *Lexer, src string, expecteds []Token) {
	tokens := tokenize(src)
	actuals := []Token{}
	for {
		actual, ok := tokens.Next()
		if !ok {
			break
		}
		actuals = append(actuals, actual)
	}

	if len(expecteds) != len(actuals) {
		t.Fatalf(`Unexpected tokenization:
Expected: %s
Actual: %s
Source:
-------
%s`, expecteds, actuals, src)
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
