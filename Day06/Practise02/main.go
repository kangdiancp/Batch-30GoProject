package main

import "fmt"

func main() {
	nums := []int{1, 5, 8, 9, 11, 13, 10, 67, 100}
	//evenNums := filterEvenNums(nums)
	evenOdd := filterOddNums(nums)

	//gaya js
	// nums.filter(x => x % 2 ==0)

	fmt.Println(evenOdd)

	primeNums := filterNums(nums, func(n int) bool {
		if n == 1 {
			return false
		}

		for i := 2; i < n/2; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	})

	evenNums := filterNums(nums, func(n int) bool {
		if n%2 == 0 {
			return true
		}
		return false
	})

	fmt.Println(evenNums)

	fmt.Println(primeNums)

}

// tujuan dari passing function parameter
// agar return dari func bisa lebih dinamis

type actionFunction func(int) bool

func isEven(n int) bool {
	return n%2 == 0
}

//bisa jg function di declare dalam parameter
func filterNums(nums []int, condition func(int) bool) []int {
	out := []int{}
	for _, v := range nums {
		if condition(v) {
			out = append(out, v)
		}
	}
	return out
}

func filterEvenNums(nums []int) []int {
	out := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			out = append(out, v)
		}
	}
	return out
}

func filterOddNums(nums []int) []int {
	out := []int{}
	for _, v := range nums {
		if v%2 != 0 {
			out = append(out, v)
		}
	}
	return out
}

func filterMultiply12Nums(nums []int) []int {
	out := []int{}
	for _, v := range nums {
		if v%12 != 0 {
			out = append(out, v)
		}
	}
	return out
}

func isOdd(n int) bool {
	return n%2 != 0
}
