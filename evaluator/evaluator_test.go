package evaluator

import (
	"programming-lang/lexer"
	"programming-lang/object"
	"programming-lang/parser"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func perform(input string) object.Object {
	ast := parser.Parse(lexer.Tokenize(input))
	return Eval(ast)
}

func TestEvalIntegerExpression(t *testing.T) {
	tdt := []struct {
		input    string
		expected int
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
	}
	for _, tc := range tdt {
		t.Run(tc.input, func(t *testing.T) {
			result := perform(tc.input)
			testInteger(t, result, tc.expected)
		})
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tdt := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"!false", true},
		{"!true", false},
		
		{"!!true", true},
		{"!!false", false},
		
		{"!5", false},
		{"!!5", true},
	}
	for _, tc := range tdt {
		t.Run(tc.input, func(t *testing.T) {
			result := perform(tc.input)
			testBoolean(t, result, tc.expected)
		})
	}
}

func testInteger(t *testing.T, ob object.Object, expected int) {
	integer, ok := ob.(*object.Integer)
	require.True(t, ok, "expected integer object, not found")
	assert.Equal(t, expected, integer.Value)
}

func testBoolean(t *testing.T, ob object.Object, expected bool) {
	boolean, ok := ob.(*object.Boolean)
	require.True(t, ok, "expected boolean object, not found")
	assert.Equal(t, expected, boolean.Value)
}