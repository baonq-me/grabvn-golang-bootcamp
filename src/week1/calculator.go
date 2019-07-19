package main

func eval(a, b float64, op func(a, b float64) (r float64, err error)) (r float64, err error) {

	return op(a, b)
}
