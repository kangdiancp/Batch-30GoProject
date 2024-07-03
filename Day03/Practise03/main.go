package main

import "fmt"

func main() {
	//1.initial value, immutable, value original
	first := 100
	fmt.Printf("[1]first:=%T [%v]\t\tAddress[%p]\n", first, first, &first)

	//2. passing by value
	val1 := incValue(first)
	fmt.Printf("[2-2]val1:=%T [%v]\t\tAddress[%p]\n", val1, val1, &val1)
	fmt.Printf("[2-3]first:=%T [%v]\t\tAddress[%p]\n", first, first, &first)

	//3. passing by reference,first : non immutable
	val2 := incPointer(&first)
	fmt.Printf("[3-2]val2:=%T [%v]\tAddress[%p]\n", val2, val2, &val2)
	fmt.Printf("[3-3]first:=%T [%v]\t\tAddress[%p]\n", first, first, &first)
}

//passing by reference
func incPointer(n *int) *int {
	fmt.Printf("[3-1]n:=%T [%v]\tAddress[%p]\n", n, n, &n)
	*n += 20
	return n
}

//passing by value
func incValue(n int) int {
	fmt.Printf("[2-1]n:=%T [%v]\t\tAddress[%p]\n", n, n, &n)
	n += 10
	return n
}
