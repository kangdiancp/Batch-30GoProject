package main

import "fmt"

func main() {
	//1. declare pointers
	var zero *int
	fmt.Printf("zero:=%T [%v]\t\tAddress[%p]\n", zero, zero, &zero)

	first := 100
	second := &first
	third := &second

	fmt.Printf("1.first:=%T [%v]\t\tAddress[%p]\n", first, first, &first)

	//2. manipulasi value variable first
	*second = 200

	fmt.Printf("2.first:=%T [%v]\t\tAddress[%p]\n", first, first, &first)
	fmt.Printf("second:=%T [%v]\tAddress[%p]\n", second, second, &second)
	fmt.Printf("third:=%T [%v]\tAddress[%p]\n", third, third, &third)
}
