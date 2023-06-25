package main

import (
	"fmt"
)

// here, phone and smartPhone are interfaces with the same method
// so any type that implements phone can be converted to a smartPhone
// and vice versa
type phone interface {
	call()
}

type smartPhone interface {
	call()
	browse()
}

// myPhone implements phone, as well as smartPhone
// since it has both the call() and browse() methods
type myPhone struct {
}

func (p myPhone) browse() {
	fmt.Println("browsing the internet")
}
func (p myPhone) call() {
	fmt.Println("ring ring")
}

func convertInterfaces() {
	var p phone = myPhone{}
	// this is a type assertion, which converts the `phone` type to a `smartPhone` type
	var sp smartPhone = p.(smartPhone)
	// now we can use methods of the `smartPhone` interface
	sp.browse()
}

type myOtherPhone struct {
}

func (p myOtherPhone) call() {
	fmt.Println("ring ring")
}

func convertIncorrectInterfaces() {
	var p phone = myOtherPhone{}
	// this will result in a runtime error
	var sp smartPhone = p.(smartPhone)
	sp.browse()
}

func main() {
	convertInterfaces()
	convertIncorrectInterfaces()
}
