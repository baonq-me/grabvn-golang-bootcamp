package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

	for scanner.Scan() {

		text := scanner.Text()

		if verify(text) {

			expr := strings.Split(strings.Replace(text, "  ", " ", -1), " ")

			// Get a, b as float64 and operator
			var a, _ = strconv.ParseFloat(expr[0], 10)
			var b, _ = strconv.ParseFloat(expr[2], 10)
			var operator = expr[1]

			// Calculate a and b
			r, e := eval(a, b, operators[operator])

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
