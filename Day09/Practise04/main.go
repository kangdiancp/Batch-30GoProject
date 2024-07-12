package main

import (
	"fmt"
	"sync"
)

var saldo = 1000

// potensi terjadi race condition
func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go withdraw(&wg, &m)
	}

	//nyuruh si main goroutine tuk nunggu
	wg.Wait()
	fmt.Println("saldo : ", saldo)

}

func withdraw(wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	saldo = saldo - 1
	m.Unlock()
}
