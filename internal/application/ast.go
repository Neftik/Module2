package application

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// ASTNode представляет узел дерева разбора.
type ASTNode struct {
	IsLeaf        bool
	Value         float64
	Operator      string
	Left, Right   *ASTNode
	TaskScheduled bool // Флаг, что для этого узла уже запланирована задача
}

// ParseAST преобразует строку с арифметическим выражением в AST.
func ParseAST(expression string) (*ASTNode, error) {
	expr := strings.ReplaceAll(expression, " ", "")
	if expr == "" {
		return nil, fmt.Errorf("empty expression")
	}
	p := &parser{input: expr, pos: 0}
	node, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	if p.pos < len(p.input) {
		return nil, fmt.Errorf("unexpected token at position %d", p.pos)
	}
	return node, nil
}

type parser struct {
	input string
	pos   int
}

func (p *parser) peek() rune {
	if p.pos < len(p.input) {
		return rune(p.input[p.pos])
	}
	return 0
}

func (p *parser) get() rune {
	ch := p.peek()
	p.pos++
	return ch
}

// parseExpression обрабатывает операции + и -.
func (p *parser) parseExpression() (*ASTNode, error) {
	node, err := p.parseTerm()
	if err != nil {
		return nil, err
	}
	for {
		ch := p.peek()
		if ch == '+' || ch == '-' {
			op := string(p.get())
			right, err := p.parseTerm()
			if err != nil {
				return nil, err
			}
			node = &ASTNode{
				IsLeaf:   false,
				Operator: op,
				Left:     node,
				Right:    right,
			}
		} else {
			break
		}
	}
	return node, nil
}

// parseTerm обрабатывает операции * и /.
func (p *parser) parseTerm() (*ASTNode, error) {
	node, err := p.parseFactor()
	if err != nil {
		return nil, err
	}
	for {
		ch := p.peek()
		if ch == '*' || ch == '/' {
			op := string(p.get())
			right, err := p.parseFactor()
			if err != nil {
				return nil, err
			}
			node = &ASTNode{
				IsLeaf:   false,
				Operator: op,
				Left:     node,
				Right:    right,
			}
		} else {
			break
		}
	}
	return node, nil
}

// parseFactor обрабатывает числа и выражения в скобках.
func (p *parser) parseFactor() (*ASTNode, error) {
	ch := p.peek()
	if ch == '(' {
		p.get() // потребляем '('
		node, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if p.peek() != ')' {
			return nil, fmt.Errorf("missing closing parenthesis")
		}
		p.get() // потребляем ')'
		return node, nil
	}
	// Обработка числа (возможно, с ведущим + или -)
	start := p.pos
	if ch == '+' || ch == '-' {
		p.get()
	}
	for {
		ch = p.peek()
		if unicode.IsDigit(ch) || ch == '.' {
			p.get()
		} else {
			break
		}
	}
	token := p.input[start:p.pos]
	if token == "" {
		return nil, fmt.Errorf("expected number at position %d", start)
	}
	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid number %s", token)
	}
	return &ASTNode{
		IsLeaf: true,
		Value:  value,
	}, nil
}
