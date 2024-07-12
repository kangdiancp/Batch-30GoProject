package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	g := runtime.GOMAXPROCS(4)
	fmt.Println("procs : ", g)
	var wg sync.WaitGroup
	wg.Add(1)

	go task("g1", &wg)
	//go task("g2", &wg)
	//go task("g3", &wg)
	fmt.Println("Main : waiting main goroutine to finish")
	//time sleep, repalce by wg.wait
	wg.Wait()
	fmt.Println("Main goroutine completed")
}

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		//go exit
		if i == 5 {
			runtime.Goexit()
		}

		fmt.Printf("%v index %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
	//will decrement counter in workgroup
	wg.Done()
}
