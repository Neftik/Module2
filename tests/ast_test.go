package tests

import (
	"fmt"
	"testing"

	"github.com/Andreyka-coder9192/calc_go/internal/application"
)

func evalAST(node *application.ASTNode) (float64, error) {
	if node.IsLeaf {
		return node.Value, nil
	}
	left, err := evalAST(node.Left)
	if err != nil {
		return 0, err
	}
	right, err := evalAST(node.Right)
	if err != nil {
		return 0, err
	}
	switch node.Operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", node.Operator)
	}
}

func TestParseASTValid(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
	}{
		{"1+2", 3},
		{"3", 3},
		{"(1+2)*3", 9},
		{"(4/2)-1", 1},
		{"2+3*4", 14}, // 2+(3*4)=14, учитывая приоритет операций
	}
	for _, tc := range tests {
		ast, err := application.ParseAST(tc.expr)
		if err != nil {
			t.Errorf("Неожиданная ошибка для выражения %s: %v", tc.expr, err)
			continue
		}
		result, err := evalAST(ast)
		if err != nil {
			t.Errorf("Ошибка вычисления AST для %s: %v", tc.expr, err)
			continue
		}
		if result != tc.expected {
			t.Errorf("Для выражения %s ожидается %f, получено %f", tc.expr, tc.expected, result)
		}
	}
}

func TestParseASTInvalid(t *testing.T) {
	invalidExprs := []string{
		"",
		"1+",
		"(1+2",
		"1++2",
		"abc",
	}
	for _, expr := range invalidExprs {
		_, err := application.ParseAST(expr)
		if err == nil {
			t.Errorf("Ожидалась ошибка для некорректного выражения %q, но ошибки не произошло", expr)
		}
	}
}
