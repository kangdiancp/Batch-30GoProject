package main

import "fmt"

func main() {
	//1. slice of string
	var slice1 []string

	//2.slice with empty state
	slice2 := []string{}

	//3.slice with length& capacity of 5
	slice3 := make([]string, 5)

	//4.slice with length& capacity of 8
	slice4 := make([]string, 5, 8)

	//5.slice with set values & length& capacity of 8
	slice5 := []string{"a", "b", "c", "d"}

	infoSlice(slice1)
	infoSlice(slice2)
	infoSlice(slice3)
	infoSlice(slice4)
	infoSlice(slice5)

	//5.1. append element in slice5
	slice5 = append(slice5, "e")
	infoSlice(slice5)

	slice5 = append(slice5, "e", "f", "g", "h")
	infoSlice(slice5)
}

func infoSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] [%p] [%s]\n", i, &slice[i], slice[i])
	}
}
