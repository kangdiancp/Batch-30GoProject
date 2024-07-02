package main

import "fmt"

func main() {
	//1. declare array default value zero
	var array [3]int
	array[0] = 15
	array[1] = 10
	array[2] = 5

	//2. declare array literally
	numbers := [5]int{1, 3, 4, 2, 1}
	fmt.Println("numbers[1]:", numbers[1])

	//3. declare array with three period
	films := [...]string{"action", "drama", "adventure"}

	//4.looping array
	for i, v := range films {
		fmt.Printf("%d : %s\n", i, v)
	}

	//5. declare array with key value
	score1 := [3]string{0: "a", 1: "b", 2: "c"}
	score2 := [3]string{"x", "y", "c"}
	fmt.Println(score1[0] == score2[0])
	fmt.Println(score1[2] == score2[2])

	//6. show address array fruits
	fruits := [5]string{"apel", "jeruk", "tomat", "durian", "mangga"}

	for i, v := range fruits {
		fmt.Printf("Values[%s]\t Address[%p] IndexElementAddr[%p]\n", v, &v, &fruits[i])
	}

}
