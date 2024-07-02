package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//1. initial array
	fruits := [5]string{"apel", "jeruk", "tomat", "durian", "mangga"}

	displayArray(fruits)

	fmt.Println(fillArrayRandom(5))

	fmt.Println(fillPrimeNumber(5))

	key := findElement(fillArrayRandom(5), 4)
	fmt.Println(key)

}

func findElement(arr [5]int, key int) int {
	for i, v := range arr {
		if v == key {
			return i
		}
	}
	return -1
}

func fillArrayRandom(n int) [5]int {
	var arr [5]int
	for i := 0; i < n; i++ {
		arr[i] = random(1, 5)
	}
	return arr
}

func fillPrimeNumber(n int) [5]int {
	var arr [5]int
	for i := 0; i < n; i++ {
		//goto, label with target
	target:
		random := random(1, 100)
		if isPrime(random) {
			arr[i] = random
		} else {
			goto target
		}

	}
	return arr
}

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			//set return & ketemu kondisi, akan keluar dari looping
			return false
		}
	}
	return true
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func displayArray(buah [5]string) {
	for i, v := range buah {
		fmt.Printf("Values[%s]\t Address[%p] IndexElementAddr[%p]\n", v, &v, &buah[i])
	}
}
