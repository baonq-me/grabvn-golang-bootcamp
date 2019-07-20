package calculator

func eval(a, b float64, op operator) (r float64, err error) {

	return op(a, b)

}
