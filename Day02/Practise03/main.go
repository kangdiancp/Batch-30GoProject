package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//1. sorting array int
	numbers := [...]int{9, 3, 65, 89, 11, 4, 7}

	//1.1. before sorting
	for _, v := range numbers {
		fmt.Printf("%d ", v)
	}

	//1.2.ready to sort
	sort.Ints(numbers[:])

	//1.3.after sorting
	fmt.Println(numbers)

	//2. sorting array string
	films := [...]string{"avenger", "saga", "madmax", "boys", "got"}
	sort.Strings(films[:])
	fmt.Println(films)

	//3. compare array
	movies := [...]string{"avenger", "saga", "Admax", "boys", "got"}

	sort.Strings(movies[:])

	fmt.Println("movies : ", movies)

	same := films == movies

	fmt.Println("same : ", same)

	rune1 := 'A'
	rune2 := 'a'

	word1 := strings.ToUpper("kecil")

	fmt.Println(rune1)
	fmt.Println(rune2)

	fmt.Println(word1)

}
