package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	/*if len(os.Args) != 4 {
		fmt.Println("Require 3 arguments !")
		return
	}

	fmt.Println(os.Args)*/

	// Can bring this to another file
	fpMap := make(map[string]func(a, b float64) (r float64, err error))
	fpMap["+"] = add
	fpMap["-"] = sub
	fpMap["*"] = mul
	fpMap["/"] = div

	//r, _ :=  eval(5, 7, fpMap["+"])
	//fmt.Println(r)

	// Regex to check expr
	// Can bring this to another file
	regex, _ := regexp.Compile(`\d+.?\d*\ +(\+|\-|\*|\/){1}\ +\d+.?\d*`)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

	for scanner.Scan() {

		text := scanner.Text()

		if regex.MatchString(text) {

			expr := strings.Split(text, " ")

			// Get a, b as float64 and operator
			var a, _ = strconv.ParseFloat(expr[0], 10)
			var b, _ = strconv.ParseFloat(expr[2], 10)
			var op = expr[1]

			// Calculate a and b
			r, e := eval(a, b, fpMap[op])

			if e == nil {
				fmt.Println(text, "=", r)
			} else {
				fmt.Println(e)
			}
		} else {
			fmt.Println("Invalid expr")
		}

		fmt.Print("> ")
	}

	return
}

/*func eval(text string) {

	expr := strings.Split(text, " ")

	var a, _ = strconv.ParseFloat(expr[0], 10)
	var b, _ = strconv.ParseFloat(expr[2], 10)
	var op = expr[1]

	switch op {
	case "+":
		fmt.Println(a, op, b, "=", a+b)
	case "-":
		fmt.Println(a, op, b, "=", a-b)
	case "*":
		fmt.Println(a, op, b, "=", a*b)
	case "/":
		fmt.Println(a, op, b, "=", a/b)
	default:
		fmt.Println("Invalid operator. Allowed operators: + - * /")
	}
}*/
