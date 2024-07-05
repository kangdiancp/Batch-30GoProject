package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkDuplicateChar("AYAM"))
}

// soal : "AYAM"
func checkDuplicateChar(word string) bool {
	chars := []rune(strings.ToLower(word))
	mapChar := make(map[rune]bool)

	for key := range chars {
		if mapChar[chars[key]] {
			return false
		} else {
			mapChar[chars[key]] = true
		}
	}

	return true
}
