package lexer

import (
	"certus/tokens"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `# This should be skipped
  # This too
  GET https://someurl.com
  `
	lexer := NewLexer(input)
	testCases := []struct {
		expectedKind    tokens.TokenKind
		expectedLiteral string
	}{
		{tokens.GET, "GET"},
		{tokens.VALUE, "https://someurl.com"},
	}
	for i, testCase := range testCases {
		currentToken := lexer.NextToken()
		if currentToken.Kind != testCase.expectedKind {
			t.Fatalf("test case no %d failed: expected %q but got %q", i, testCase.expectedKind, currentToken.Kind)
		}
		if currentToken.Literal != testCase.expectedLiteral {
			t.Fatalf("test case no %d failed: expected %q but got %q", i, testCase.expectedLiteral, currentToken.Literal)
		}
	}
}
