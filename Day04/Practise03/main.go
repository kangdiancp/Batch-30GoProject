package main

import (
	"fmt"
)

func main() {
	//1. init map string
	products := make(map[string]float64, 10)
	products["tv"] = 1_500_000.98
	products["book"] = 56_908.98
	products["mousepad"] = 8500.98

	//1.1. update value
	products["books"] = 150_000.98

	//1.3. delete element
	delete(products, "books")

	//2. loop
	for key, v := range products {
		fmt.Printf("key:[%s] value:[%.2f]\n", key, v)
	}

}
