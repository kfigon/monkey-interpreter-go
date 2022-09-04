package lexerbybook

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLexerByTheBook(t *testing.T) {
	input := `let asd = 123;
for(i=1;i<3; i++){}`

	expected := []llexToken{
		{"KEYWORD", "let"},
		{"IDENTIFIER", "asd"},
		{"ASSIGNMENT", "="},
		{"NUMBER", "123"},
		{"SEMICOLON", ";"},
		{"KEYWORD", "for"},
		{"OPEN_BRACE", "("},
		{"IDENTIFIER", "i"},
		{"ASSIGNMENT", "="},
		{"NUMBER", "1"},
		{"SEMICOLON", ";"},
		{"IDENTIFIER", "i"},
		{"OPERATOR", "<"},
		{"NUMBER", "3"},
		{"SEMICOLON", ";"},
		{"IDENTIFIER", "i"},
		{"OPERATOR", "++"},
		{"CLOSE_BRACE", ")"},
		{"OPEN_CBRACE", "{"},
		{"CLOSE_CBRACE", "}"},
	}

	llex := newLlexer(input)
	var got []llexToken
	for token := llex.emit(); token.token != "EOF"; token = llex.emit() {
		got = append(got, token)
	}

	assert.Equal(t, expected, got)
}

func TestLexerByTheBookSimpleIdentifier(t *testing.T) {
	l := newLlexer(`let;`)
	to := l.emit()
	semi := l.emit()

	assert.Equal(t, "EOF", l.emit().token)
	assert.Equal(t, "KEYWORD", to.token)
	assert.Equal(t, "let", to.lexeme)
	
	assert.Equal(t, "SEMICOLON", semi.token)
	assert.Equal(t, ";", semi.lexeme)
}