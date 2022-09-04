package parser

import (
	"fmt"
	"programming-lang/lexer"
	"strconv"
)

type ExpressionNode interface {
	Node
	evaluateExpression()
}

type IntegerLiteralExpression struct {
	Value int
}

func (ile *IntegerLiteralExpression) TokenLiteral() string {
	return strconv.Itoa(ile.Value)
}
func (ile *IntegerLiteralExpression) String() string {
	return strconv.Itoa(ile.Value)
}
func (ile *IntegerLiteralExpression) evaluateExpression() {}

type IdentifierExpression struct {
	Name string
}

func (ide *IdentifierExpression) TokenLiteral() string {
	return ide.Name
}

func (ide *IdentifierExpression) String() string {
	return ide.Name
}

func (ide *IdentifierExpression) evaluateExpression() {}

type PrefixExpression struct {
	Operator string
	Right    ExpressionNode
}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Operator
}

func (p *PrefixExpression) String() string {
	return "(" + p.Operator + p.Right.String() + ")"
}

func (p *PrefixExpression) evaluateExpression() {}

type InfixExpression struct {
	Operator string
	Left     ExpressionNode
	Right    ExpressionNode
}

func (i *InfixExpression) TokenLiteral() string {
	return i.Operator
}
func (i *InfixExpression) String() string {
	return "("+ i.Left.String() + i.Operator + i.Right.String() +")"
}

func (i *InfixExpression) evaluateExpression() {}

type BooleanExpression struct {
	Value    bool
}

func (b *BooleanExpression) TokenLiteral() string {
	return b.String()
}
func (b *BooleanExpression) String() string {
	return strconv.FormatBool(b.Value)
}

func (b *BooleanExpression) evaluateExpression() {}

type IfExpression struct {
	Condition ExpressionNode
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfExpression) TokenLiteral() string {
	return "if"
}

func (i *IfExpression) String() string {
	out := "if" + i.Condition.String() + " " + i.Consequence.String()
	if i.Alternative != nil {
		out += " else " + i.Alternative.String()
	}
	return out
}

func (i *IfExpression) evaluateExpression() {}

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

func (p *parser) parseExpression(predescense int) ExpressionNode {
	tok := p.currentToken

	var left ExpressionNode
	if bang(tok) || minus(tok) {
		 left = p.parsePrefixExpression()
	} else if isNumberLiteral(tok){
		left = p.parseIntegerLiteralExpression()
	} else if isBoolean(tok) {
		left = p.parseBooleanExpression()
	} else if isIdentifier(tok) {
		left = p.parseIdentifierExpression()
	} else if isOpeningParent(tok) {
		left = p.parseGroupedExpression()
	}else if ifKeyword(tok){
		left = p.parseIfExpression()
	} else {
		p.addError(fmt.Errorf("no prefix parsing function for token %s", tok.Lexeme))
		return nil
	}


	// infix - left associated
	for !isSemicolon(p.nextToken) && predescense < tokensPredescense(p.nextToken) {

		if equals(p.nextToken) || 
			notEquals(p.nextToken) || 
			lessThan(p.nextToken) || 
			lessEqThan(p.nextToken) || 
			greaterEqThan(p.nextToken) || 
			greaterThan(p.nextToken) || 
			plus(p.nextToken) || 
			minus(p.nextToken) || 
			product(p.nextToken) || 
			divide(p.nextToken) {
			
			p.advanceToken()
			left = p.parseInfixExpression(left)
		} else {
			return left
		}
	}

	return left
}

func tokensPredescense(tok lexer.Token) int {
	switch {
	case equals(tok):
		return EQUALS
	case notEquals(tok):
		return EQUALS

	case lessThan(tok):
		return LESSGREATER
	case lessEqThan(tok):
		return LESSGREATER
	case greaterEqThan(tok):
		return LESSGREATER
	case greaterThan(tok):
		return LESSGREATER

	case plus(tok):
		return SUM
	case minus(tok):
		return SUM

	case product(tok):
		return PRODUCT
	case divide(tok):
		return PRODUCT
	default:
		return LOWEST
	}
}

func (p *parser) parseIntegerLiteralExpression() ExpressionNode {
	tok := p.currentToken
	v, err := strconv.Atoi(tok.Lexeme)
	if err != nil {
		p.addError(fmt.Errorf("int literal expression error - error in parsing integer literal in: %v", tok.Lexeme))
		return nil
	}
	return &IntegerLiteralExpression{Value: v}
}

func (p *parser) parseIdentifierExpression() ExpressionNode {
	identifierToken := p.currentToken
	return &IdentifierExpression{Name: identifierToken.Lexeme}
}

func (p *parser) parseBooleanExpression() ExpressionNode {
	v, err := strconv.ParseBool(p.currentToken.Lexeme)
	if err != nil {
		p.addError(fmt.Errorf("boolean literal expression error: %v, token: %v", err, p.currentToken))
		return nil
	}
	return &BooleanExpression{v}
}


func (p *parser) parsePrefixExpression() ExpressionNode {
	operator := p.currentToken
	p.advanceToken()
	return &PrefixExpression{Operator: operator.Lexeme, Right: p.parseExpression(PREFIX)}
}

func (p *parser) parseInfixExpression(left ExpressionNode) ExpressionNode {
	out := &InfixExpression{
		Operator: p.currentToken.Lexeme,
		Left:     left,
	}
	pred := tokensPredescense(p.currentToken)
	p.advanceToken()
	out.Right = p.parseExpression(pred)
	return out
}

func (p *parser) parseGroupedExpression() ExpressionNode {
	p.advanceToken()

	out := p.parseExpression(LOWEST)
	if !isClosingParent(p.nextToken) {
		p.addError(fmt.Errorf("grouped expression error - missing closing brace, got %v", p.nextToken.Lexeme))
		return nil
	}
	p.advanceToken()

	return out
}

func (p *parser) parseIfExpression() ExpressionNode {
	if !isOpeningParent(p.nextToken) {
		p.addError(fmt.Errorf("if expression error - missing opening brace, got %v", p.nextToken.Lexeme))
		return nil
	}
	p.advanceToken()

	out := &IfExpression{}
	out.Condition = p.parseExpression(LOWEST)
	
	if !isClosingParent(p.currentToken) {
		p.addError(fmt.Errorf("if expression error - missing closing brace, got %v", p.currentToken.Lexeme))
		return nil
	}
	p.advanceToken()

	if !isOpeningCurly(p.currentToken) {
		p.addError(fmt.Errorf("if expression error - missing opening curly brace, got %v", p.currentToken.Lexeme))
		return nil
	} 

	out.Consequence = p.parseBlockStatement()
	
	if elseKeyword(p.nextToken) {
		p.advanceToken()
		if !isOpeningCurly(p.nextToken) {
			p.addError(fmt.Errorf("else expression error - missing opening curly brace, got %v", p.nextToken.Lexeme))
			return nil
		}
		p.advanceToken()
		out.Alternative = p.parseBlockStatement()
	}

	return out
}