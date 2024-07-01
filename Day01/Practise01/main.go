package main

import "fmt"

func main() {
	//okYes(15)
	sum := sumOk(15)
	fmt.Println("sum : ", sum)
	fmt.Println("direct sum : ", sumOk(35))

	fmt.Println("count ok : ", countOk(35))
}

// 1. okYes
func okYes(n int) {
	for i := 0; i < n; i++ {
		//fmt.Println("i : ", i)
		if i%3 == 0 && i%4 == 0 {
			fmt.Println("OKYes")
		} else if i%3 == 0 {
			fmt.Println("OK")
		} else if i%4 == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println(i)
		}
	}
}

//2. sumOK
func sumOk(n int) int {
	//short variable, always initial value
	sum := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 && i%4 != 0 {
			//sum = sum +i
			sum += i
		}
	}
	return sum
}

//3. count ok
func countOk(n int) int {
	//short variable, always initial value
	count := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 && i%4 != 0 {
			//sum = sum +i
			count++
		}
	}
	return count
}
