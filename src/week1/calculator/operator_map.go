package calculator

var (
	operators = map[string]operator{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
)
