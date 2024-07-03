package main

import "fmt"

func main() {
	//1. value type/primitive type
	var a int
	var b string
	var c float64
	var d bool

	// %T:data type, %v:value, %p:address memory
	fmt.Printf("a:=%T [%v]\tAddress[%p]\n", a, a, &a)
	fmt.Printf("b:=%T [%v]\tAddress[%p]\n", b, b, &b)
	fmt.Printf("c:=%T [%v]\tAddress[%p]\n", c, c, &c)
	fmt.Printf("d:=%T [%v]\tAddress[%p]\n\n", d, d, &d)

	//2. short variable
	aa := 10      // int [10]
	bb := "hello" // string [hello]
	cc := 3.14159 // float64 [3.14159]
	dd := true    // bool [true]
	fmt.Printf("aa:=%T [%v]\t\tAddress[%p]\n", aa, aa, &aa)
	fmt.Printf("bb:=%T [%v]\tAddress[%p]\n", bb, bb, &bb)
	fmt.Printf("cc:=%T [%v]\tAddress[%p]\n", cc, cc, &cc)
	fmt.Printf("dd:=%T [%v]\t\tAddress[%p]\n\n", dd, dd, &dd)
}
