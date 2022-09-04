package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		desc           string
		input          string
		expectedTokens []Token
	}{
		{
			desc:  "simple case with whitespace",
			input: ` i;`,
			expectedTokens: []Token{
				{Identifier, "i"},
				{Semicolon, ";"},
				{EOF,""},
			},
		},
		{
			desc: "complex case 1",
			input: `if (i==j) els = 654.1;
	else els=123;`,
			expectedTokens: []Token{
				{Keyword, "if"},
				{OpenParam, "("},
				{Identifier, "i"},
				{Operator, "=="},
				{Identifier, "j"},
				{CloseParam, ")"},
				{Identifier, "els"},
				{Assignment, "="},
				{Number, "654.1"},
				{Semicolon, ";"},
				{Keyword, "else"},
				{Identifier, "els"},
				{Assignment, "="},
				{Number, "123"},
				{Semicolon, ";"},
				{EOF,""},
			},
		},
		{
			desc: "case with math operators",
			input: `abc=123 / 2*1-12
	x=3+2;`,
			expectedTokens: []Token{
				{Identifier, "abc"},
				{Assignment, "="},
				{Number, "123"},
				{Operator, "/"},
				{Number, "2"},
				{Operator, "*"},
				{Number, "1"},
				{Operator, "-"},
				{Number, "12"},
				{Identifier, "x"},
				{Assignment, "="},
				{Number, "3"},
				{Operator, "+"},
				{Number, "2"},
				{Semicolon, ";"},
				{EOF,""},
			},
		},
		{
			desc:  "for loop",
			input: `for(i=0;i<3;i++){`,
			expectedTokens: []Token{
				{Keyword, "for"},
				{OpenParam, "("},
				{Identifier, "i"},
				{Assignment, "="},
				{Number, "0"},
				{Semicolon, ";"},
				{Identifier, "i"},
				{Operator, "<"},
				{Number, "3"},
				{Semicolon, ";"},
				{Identifier, "i"},
				{Operator, "++"},
				{CloseParam, ")"},
				{OpenParam, "{"},
				{EOF,""},
			},
		},
		{
			desc:  "var statement",
			input: `var foo = 123 != 3;`,
			expectedTokens: []Token{
				{Class: Keyword, Lexeme: "var"},
				{Class: Identifier, Lexeme: "foo"},
				{Class: Assignment, Lexeme: "="},
				{Class: Number, Lexeme: "123"},
				{Class: Operator, Lexeme: "!="},
				{Class: Number, Lexeme: "3"},
				{Class: Semicolon, Lexeme: ";"},
				{EOF,""},
			},
		},
		{
			desc:  "function declaration",
			input: `fn asd(){}`,
			expectedTokens: []Token{
				{Class: Keyword, Lexeme: "fn"},
				{Class: Identifier, Lexeme: "asd"},
				{Class: OpenParam, Lexeme: "("},
				{Class: CloseParam, Lexeme: ")"},
				{Class: OpenParam, Lexeme: "{"},
				{Class: CloseParam, Lexeme: "}"},
				{EOF,""},
			},
		},
		{
			desc:  "prefix operator",
			input: `var foo = -5`,
			expectedTokens: []Token{
				{Class: Keyword, Lexeme: "var"},
				{Class: Identifier, Lexeme: "foo"},
				{Class: Assignment, Lexeme: "="},
				{Class: Operator, Lexeme: "-"},
				{Class: Number, Lexeme: "5"},
				{EOF,""},
			},
		},
		{
			desc:  "infix operators",
			input: `5 == 3 > 1 >= 1 < <= != -- / *`,
			expectedTokens: []Token{
				{Class: Number, Lexeme: "5"},
				{Class: Operator, Lexeme: "=="},
				{Class: Number, Lexeme: "3"},
				{Class: Operator, Lexeme: ">"},
				{Class: Number, Lexeme: "1"},
				{Class: Operator, Lexeme: ">="},
				{Class: Number, Lexeme: "1"},
				{Class: Operator, Lexeme: "<"},
				{Class: Operator, Lexeme: "<="},
				{Class: Operator, Lexeme: "!="},
				{Class: Operator, Lexeme: "--"},
				{Class: Operator, Lexeme: "/"},
				{Class: Operator, Lexeme: "*"},
				{EOF,""},
			},
		},
		{
			desc:  "boolean",
			input: `true false falset truet true;`,
			expectedTokens: []Token{
				{Class: Boolean, Lexeme: "true"},
				{Class: Boolean, Lexeme: "false"},
				{Class: Identifier, Lexeme: "falset"},
				{Class: Identifier, Lexeme: "truet"},
				{Class: Boolean, Lexeme: "true"},
				{Class: Semicolon, Lexeme: ";"},
				{EOF,""},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Tokenize(tC.input)
			assert.Equal(t, tC.expectedTokens, got)
		})
	}
}
