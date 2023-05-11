package ast

import (
	"fmt"
	"testing"
)

func TestLexerString(t *testing.T) {
	input := `"This is a string"`
	l := NewLexer(input)
	tok := l.NextToken()
	if tok.Type != STRING {
		t.Errorf("Expected token type STRING, got %s", tok.Type)
	}
	if tok.Value != "This is a string" {
		t.Errorf("Expected token value 'This is a string', got %s", tok.Value)
	}
}

func TestLexerMultipleStrings(t *testing.T) {
	input := `"This is a string""This is another string""This is a third string"`
	l := NewLexer(input)
	tokens := []Token{}
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
		tokens = append(tokens, tok)
	}

	if len(tokens) != 3 {
		t.Errorf("Expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0].Type != STRING {
		t.Errorf("Expected token type STRING, got %s", tokens[0].Type)
	}
	if tokens[1].Type != STRING {
		t.Errorf("Expected token type STRING, got %s", tokens[1].Type)
	}
	if tokens[2].Type != STRING {
		t.Errorf("Expected token type STRING, got %s", tokens[2].Type)
	}

	if tokens[0].Value != "This is a string" {
		t.Errorf("Expected token value 'This is a string', got %s", tokens[0].Value)
	}
	if tokens[1].Value != "This is another string" {
		t.Errorf("Expected token value 'This is another string', got %s", tokens[1].Value)
	}
	if tokens[2].Value != "This is a third string" {
		t.Errorf("Expected token value 'This is a third string', got %s", tokens[2].Value)
	}

}

func TestLexerNumber(t *testing.T) {
	input := `1234567890`
	l := NewLexer(input)
	tokens := []Token{}
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
		tokens = append(tokens, tok)
	}

	if len(tokens) != 1 {
		t.Errorf("Expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Type != NUMBER {
		t.Errorf("Expected token type NUMBER, got %s", tokens[0].Type)
	}
	if tokens[0].Value != "1234567890" {
		t.Errorf("Expected token value '1234567890', got %s", tokens[0].Value)
	}
}

func TestLexerMultipleNumbers(t *testing.T) {
	input := `1234567890 1234567890 1234567890`
	l := NewLexer(input)
	tokens := []Token{}
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
		tokens = append(tokens, tok)
	}

	if len(tokens) != 3 {
		t.Errorf("Expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0].Type != NUMBER {
		t.Errorf("Expected token type NUMBER, got %s", tokens[0].Type)
	}
	if tokens[1].Type != NUMBER {
		t.Errorf("Expected token type NUMBER, got %s", tokens[1].Type)
	}
	if tokens[2].Type != NUMBER {
		t.Errorf("Expected token type NUMBER, got %s", tokens[2].Type)
	}

	if tokens[0].Value != "1234567890" {
		t.Errorf("Expected token value '1234567890', got %s", tokens[0].Value)
	}
	if tokens[1].Value != "1234567890" {
		t.Errorf("Expected token value '1234567890', got %s", tokens[1].Value)
	}
	if tokens[2].Value != "1234567890" {
		t.Errorf("Expected token value '1234567890', got %s", tokens[2].Value)
	}
}
