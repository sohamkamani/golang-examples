package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type child struct {
	name string
	age  int
}

type pet struct {
	name string
}

func convertStructs() {
	bob := person{
		name: "bob",
		age:  15,
	}
	babyBob := child(bob)
	fmt.Println(bob, babyBob)
	// output:
	// {bob 15} {bob 15}
}

func convertBasicTypes() {
	// ints can be converted to floats
	var n int = 5
	fmt.Println(float64(n))
	// output:
	// 5

	// floats can be converted to ints
	// but the fraction is eliminated
	var pi float64 = 3.1412
	fmt.Println(int(pi))
	// output:
	// 3

	// strings can be converted to byte arrays
	greeting := "hello world"
	fmt.Println([]byte(greeting))
	// output:
	// [104 101 108 108 111 32 119 111 114 108 100]
}

// here, phone and smartPhone are interfaces with the same method
// so any type that implements phone can be converted to a smartPhone
// and vice versa
type phone interface {
	call()
}

type smartPhone interface {
	call()
}

// myPhone implements phone, as well as smartPhone
type myPhone struct {
}

func (p myPhone) call() {
	fmt.Println("ring ring")
}

func convertInterfaces() {
	var p phone = myPhone{}
	var sp smartPhone = p
	sp.call()
	// output:
	// ring ring
}

func main() {
	convertStructs()
	convertBasicTypes()
	convertInterfaces()
}
