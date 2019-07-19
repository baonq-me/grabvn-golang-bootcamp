package main

var (
	fpMap = initOperatorMap()
)

func initOperatorMap() map[string]func(a, b float64) (r float64, err error) {

	fpMap := make(map[string]func(a, b float64) (r float64, err error))
	fpMap["+"] = add
	fpMap["-"] = sub
	fpMap["*"] = mul
	fpMap["/"] = div

	return fpMap
}
