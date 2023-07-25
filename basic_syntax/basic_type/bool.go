package main

import "fmt"

func Bool() {
	var a bool = true
	var b bool = false
	var c bool = a || b
	fmt.Printf("a || b=: %s\n", c)
	var d bool = a && b
	fmt.Printf("a && b=: %s\n", d)
	var e bool = !a
	fmt.Printf("!a=: %s\n", e)
	var x bool = !(a && b)
	fmt.Printf("!(a && b)=: %s\n", x)
	var y bool = !(a || b)
	fmt.Printf("!(a || b)=: %s\n", y)
}
