package evaluator

import (
	"fmt"
	"turkey-lang/ast"
	"turkey-lang/object"
	"turkey-lang/token"
)

func show(node ast.Node, env *object.Environment) object.Object {
	node = evalPlayCalls(node, env)
	return &object.Show{Node: node}
}

func evalPlayCalls(shown ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(shown, func(node ast.Node) ast.Node {
		if !isPlayCall(node) {
			return node
		}

		call, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}
		if len(call.Arguments) != 1 {
			return node
		}

		played := Eval(call.Arguments[0], env)
		return convertObjectToASTNode(played)
	})
}

func convertObjectToASTNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		t := token.Token{
			Type:    token.INT,
			Literal: fmt.Sprintf("%d", obj.Value),
		}
		return &ast.IntegerLiteral{Token: t, Value: obj.Value}
	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{Type: token.TRUE, Literal: "true"}
		} else {
			t = token.Token{Type: token.FALSE, Literal: "false"}
		}
		return &ast.Boolean{Token: t, Value: obj.Value}
	default:
		return nil
	}
}

func isPlayCall(node ast.Node) bool {
	callExpression, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}

	return callExpression.Function.TokenLiteral() == "play"
}
