package main

import "fmt"

func main() {
	/* n := 3456
	count := 0

	sisa := n / 10
	fmt.Println("sisa1 : ", sisa)
	count++

	sisa = sisa / 10
	fmt.Println("sisa2 : ", sisa)
	count++

	sisa = sisa / 10
	fmt.Println("sisa3 : ", sisa)
	count++

	fmt.Println("countdigit :", count) */

	fmt.Println("countDigit : ", countDigit(3456))

	fmt.Println("countDigit : ", countDigit(345342346))

	fmt.Println("countDigit : ", countDigit(34563434234234))

	fmt.Println("isPrime :", isPrime(7))
	fmt.Println("isPrime :", isPrime(8))

	fibonacci(10)
}

func countDigit(n int) int {
	count := 0
	sisa := n
	for {
		sisa = sisa / 10
		count++
		if sisa == 0 {
			break
		}
	}
	return count
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

func fibonacci(n int) {
	first := 0
	second := 1
	for i := 0; i < n; i++ {
		fibo := first + second
		first = second
		second = fibo
		fmt.Printf("%d ", fibo)
	}
}
