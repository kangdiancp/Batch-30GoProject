package main

import (
	fmr "day06/practise01/mypackage/formatter" //overriding with alias
	"day06/practise01/mypackage/math"
	"fmt"
)

func main() {
	n := math.Multiply(3, 4)
	fmt.Println(n)

	out := fmr.FormatNum(n)
	fmt.Println(out)

	//handling error
	n, error := math.MultiplyEven(13)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(n)

}
