package main

import "fmt"

func main() {
	/* printBox(5)
	printBoxNumber(5)
	printBoxHollow(5)
	printTriangle(5)
	printTriangleNumber(5)
	printTriangle2(5) */
	diagonal(5)
}

func printBox(n int) {
	//cetak start horizontal
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

func printBoxNumber(n int) {
	//cetak start horizontal
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			fmt.Printf("%d ", row+col+1)
		}
		fmt.Println("")
	}
}

func printBoxHollow(n int) {
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row == 0 || row == n-1 || col == 0 || col == n-1 {
				fmt.Printf("%s ", "*")
			} else {
				fmt.Printf("%s ", " ")
			}
		}
		fmt.Println("")
	}
}

func printTriangle(n int) {
	for row := 1; row <= n; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%s ", "*")
		}
		fmt.Println("")
	}
}

func printTriangleNumber(n int) {
	for row := 1; row <= n; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%d ", row+col)
		}
		fmt.Println("")
	}
}

func printTriangle2(n int) {
	for row := 1; row <= n; row++ {
		for col := n; col >= row; col-- {
			fmt.Printf("%s ", "*")
		}
		fmt.Println("")
	}
}

func diagonal(n int) {
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row == col {
				fmt.Printf("%d ", row+col+1)
			} else {
				fmt.Printf("%s", "*")
			}
		}
		fmt.Println("")
	}
}
