package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//1.fill matrix array manually
	var matrix1 [3][3]int
	matrix1[0][0] = 10
	matrix1[0][1] = 15
	matrix1[0][2] = 15

	//2.matrix with short variable
	matrix2 := [3][3]int{
		{0, 1, 2},
		{0, 1, 2},
		{0, 1, 5},
	}

	displayMatrix(matrix2)
	fmt.Println("")

	var matrix3 [3][3]int
	matrix3 = fillMatrix(matrix3)
	displayMatrix(matrix3)

	//3. fill matrix array with random number
	var matrix [5][5]int
	matrix = fillArrayWithRandom(matrix)
	displayMatrixFive(matrix)

	//4. fill matrix array with prime number
	matrixPrime := fillArrayWithPrime(matrix)
	displayMatrixFive(matrixPrime)
}

func displayMatrixFive(matrix [5][5]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println("")
	}
}

func displayMatrix(matrix [3][3]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println("")
	}
}

func fillMatrix(matrix [3][3]int) [3][3]int {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = row + col
		}
	}
	return matrix
}

func fillArrayWithRandom(matrix [5][5]int) [5][5]int {

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = rand.Intn(11)
		}
	}
	return matrix
}

func fillArrayWithPrime(matrix [5][5]int) [5][5]int {

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
		target:
			random := rand.Intn(row + 5)

			if isPrime(random) {
				matrix[row][col] = random
			} else {
				goto target
			}

		}
	}
	return matrix
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
