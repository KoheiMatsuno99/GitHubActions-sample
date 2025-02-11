package sample

import "errors"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot devide by zero")
	}
	return a / b, nil
}
