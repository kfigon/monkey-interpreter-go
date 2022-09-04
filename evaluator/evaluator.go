package evaluator

import (
	"programming-lang/object"
	"programming-lang/parser"
)

var (
	TRUE_VAL = &object.Boolean{Value: true}
	FALSE_VAL = &object.Boolean{Value: false}
	NULL_VAL = &object.Null{}
)


func Eval(node parser.Node) object.Object {
	switch n := node.(type) {
	case *parser.Program:
		return evalStatemnets(n)
	case *parser.IntegerLiteralExpression:
		return &object.Integer{Value: n.Value}
	case *parser.BooleanExpression:
		return evalBoolean(n)
	case *parser.ExpressionStatementNode:
		return Eval(n.Value)
	case *parser.PrefixExpression:
		return evalPrefix(n)
	}
	return nil
}

func evalStatemnets(node *parser.Program) object.Object {
	var out object.Object
	for _, v := range node.Statements {
		out = Eval(v)
	}
	return out
}

func evalBoolean(node *parser.BooleanExpression) object.Object {
	switch node.Value {
	case true: return TRUE_VAL
	case false: return FALSE_VAL
	}
	return nil
}

func evalPrefix(node *parser.PrefixExpression) object.Object {
	right := Eval(node.Right)
	if node.Operator == "!" {
		switch right {
		case TRUE_VAL: return FALSE_VAL
		case FALSE_VAL: return TRUE_VAL
		case NULL_VAL: return TRUE_VAL
		default: return FALSE_VAL
		}
	} else if node.Operator == "-" && right.Type() == object.INTEGER {
		v := right.(*object.Integer).Value
		return &object.Integer{Value: -v}
	}
	return NULL_VAL
}