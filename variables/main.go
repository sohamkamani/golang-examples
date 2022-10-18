package main

import "fmt"

func main() {
	// declaration with type
	var i1 int = 5

	// declaration with inferred type
	var i2 = 5

	// declaration without assignment
	var i3 int

	// This is equivalent to
	// var i4 int = 5
	i4 := 5

	// this code will not compile
	// var i int
	// i := 5

	a, b := 1, 2
	// We have declared and initialized the variables `a` and `b` with the values of 1 and 2 respectively

	// We can even use this to assign values later
	var c int
	var d int
	c, d = 5, 3

	// These are all private variables, and cannot be accessed outside of this package
	var i5 int
	count := 5

	// These are public variables
	var Name string
	Age := 24

	var (
		// we can declare variables together in a single `var` block
		name = "John Doe"
		age  = 30
		// we can assign multiple values at once
		isLorem, isIpsum = false, true
		// we can declare variables without assigning values
		info string
	)

	fmt.Println(
		i1,
		i2,
		i3,
		i4,
		a,
		b,
		c,
		d,
		i5,
		count,
		Name,
		Age,
		name,
		age,
		isLorem,
		isIpsum,
		info,
	)
}
