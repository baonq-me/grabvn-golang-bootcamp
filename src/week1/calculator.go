package main

import (
	"errors"
)

func eval(a, b float64, op func(a, b float64) (r float64, err error)) (r float64, err error) {

	return op(a, b)
}

func add(a, b float64) (r float64, err error) {
	return a + b, nil
}

func sub(a, b float64) (r float64, err error) {
	return a - b, nil
}

func mul(a, b float64) (r float64, err error) {
	return a * b, nil
}

func div(a, b float64) (r float64, err error) {
	if b-0.0000000001 < 0 {
		err = errors.New("Div by zero")
		return
	}

	return a / b, nil
}