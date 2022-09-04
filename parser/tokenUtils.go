package parser

import "programming-lang/lexer"

func isVarKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "var"
}

func isAssignmentOperator(token lexer.Token) bool {
	return token.Class == lexer.Assignment && token.Lexeme == "="
}

func isSemicolon(token lexer.Token) bool {
	return token.Class == lexer.Semicolon && token.Lexeme == ";"
}

func isNumberLiteral(token lexer.Token) bool {
	return token.Class == lexer.Number
}

func isIdentifier(token lexer.Token) bool {
	return token.Class == lexer.Identifier
}

func isBoolean(token lexer.Token) bool {
	return token.Class == lexer.Boolean
}

func isOpeningParent(token lexer.Token) bool {
	return token.Class == lexer.OpenParam && token.Lexeme == "("
}

func isClosingParent(token lexer.Token) bool {
	return token.Class == lexer.CloseParam && token.Lexeme == ")"
}

func isOpeningCurly(token lexer.Token) bool {
	return token.Class == lexer.OpenParam && token.Lexeme == "{"
}

func isClosingCurly(token lexer.Token) bool {
	return token.Class == lexer.CloseParam && token.Lexeme == "}"
}

func isReturnKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "return"
}

func ifKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "if"
}

func elseKeyword(token lexer.Token) bool {
	return token.Class == lexer.Keyword && token.Lexeme == "else"
}

func eof(token lexer.Token) bool {
	return token.Class == lexer.EOF
}

func bang(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "!"
}

func minus(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "-"
}

func product(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "*"
}

func plus(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "+"
}

func divide(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "/"
}

func equals(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "=="
}

func notEquals(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "!="
}

func lessThan(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "<"
}

func lessEqThan(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == "<="
}

func greaterThan(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == ">"
}

func greaterEqThan(token lexer.Token) bool {
	return token.Class == lexer.Operator && token.Lexeme == ">="
}
