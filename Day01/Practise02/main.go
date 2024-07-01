package main

import "fmt"

func main() {
	hero := "tanjiro"

	// string to byte
	for i := 0; i < len(hero); i++ {
		fmt.Print(hero[i], " ")
	}

	fmt.Println("")

	//2. byte to string
	for i := 0; i < len(hero); i++ {
		fmt.Print(string(hero[i]), " ")
	}

	fmt.Println("")

	//3. using range
	for index, v := range hero {
		fmt.Printf("%d:%v ", index, v)
	}

	//4. string to rune
	runeHero := []rune(hero)
	for i := 0; i < len(runeHero); i++ {
		fmt.Print("i : ", i)
	}

	//5. count vowel
	fmt.Println("")
	movies := "madmax furiosa"
	fmt.Println("count vowel :", countVowel(movies))
	x := 'a'
	fmt.Println("x:", x)

	//6. count vowel with multiple return
	vocal, consonant := countVowel2(movies)
	fmt.Printf("vocal : %d, consonant: %d", vocal, consonant)

}

func countVowel(word string) int {
	result := 0
	for i := 0; i < len(word); i++ {
		if word[i] == 'a' || word[i] == 'i' || word[i] == 'u' || word[i] == 'e' ||
			word[i] == 'o' {
			//result = result+1
			result++
		}
	}
	return result
}

//multiple return
func countVowel2(word string) (int, int) {
	//short variable
	vocal, consonant := 0, 0
	// default zero
	//var vocal, consonant int

	for i := 0; i < len(word); i++ {
		if word[i] == 'a' || word[i] == 'i' || word[i] == 'u' || word[i] == 'e' ||
			word[i] == 'o' {
			vocal++
		} else if word[i] != ' ' {
			consonant++
		}
	}
	return vocal, consonant

}
