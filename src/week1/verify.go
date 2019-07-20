package main

import "regexp"

var (
	regex, _ = regexp.Compile(`^\-?\d+.?\d*\ +(\+|\-|\*|\/){1}\ \-?\d+.?\d*$`)
)

func verify(s string) bool {

	return regex.MatchString(s)

}
