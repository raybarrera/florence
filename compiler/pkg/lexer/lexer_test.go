package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `
		program main

		import {
			"async"
			"foo"
		}

		type rules trait {
			apply()
		}

		type game struct {
			int fps
		}


		fn main(...string argv) {
		}
	`

	expectedTokens := []Token{
		{Type: TokenIdentifier, Value: "program"},
		{Type: TokenIdentifier, Value: "main"},
		{Type: TokenKeyword, Value: "import"},
		{Type: TokenOpenBrace, Value: "{"},
		{Type: TokenString, Value: "async"},
		{Type: TokenString, Value: "foo"},
		{Type: TokenCloseBrace, Value: "}"},
		{Type: TokenKeyword, Value: "type"},
		{Type: TokenIdentifier, Value: "rules"},
		{Type: TokenIdentifier, Value: "trait"},
		{Type: TokenOpenParen, Value: "("},
		{Type: TokenCloseParen, Value: ")"},
		{Type: TokenOpenBrace, Value: "{"},
		{Type: TokenIdentifier, Value: "apply"},
		{Type: TokenOpenParen, Value: "("},
		{Type: TokenCloseParen, Value: ")"},
		{Type: TokenCloseBrace, Value: "}"},
		// ...rest of the expected tokens...
		{Type: TokenEof, Value: ""},
	}

	lexer := NewLexer(input)

	for i, expectedToken := range expectedTokens {
		token := lexer.NextToken()

		if token.Type != expectedToken.Type {
			t.Errorf("Token type mismatch at index %d. Expected: %v, got: %v", i, expectedToken.Type, token.Type)
		}

		if token.Value != expectedToken.Value {
			t.Errorf("Token value mismatch at index %d. Expected: %s, got: %s", i, expectedToken.Value, token.Value)
		}
	}

	fmt.Println("Lexer test passed!")
}
