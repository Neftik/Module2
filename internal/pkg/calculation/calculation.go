package calculation

import (
	"fmt"
)

// Calc – старая функция вычисления выражения (не используется в распределённом варианте).
func Calc(expression string) (float64, error) {
	return 0, fmt.Errorf("not implemented")
}

// Compute выполняет бинарную операцию над двумя числами.
func Compute(operation string, a, b float64) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", operation)
	}
}
