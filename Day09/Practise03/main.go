package main

import (
	"fmt"
	"sync"
)

var saldo = 1000

// potensi terjadi race condition
func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go withdraw(&wg)
	}

	//nyuruh si main goroutine tuk nunggu
	wg.Wait()
	fmt.Println("saldo : ", saldo)

}

func withdraw(wg *sync.WaitGroup) {
	defer wg.Done()
	saldo = saldo - 1 //990
}
