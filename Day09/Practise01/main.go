package main

import (
	"fmt"
	"time"
)

func main() {

	defer fmt.Println("siap sajikan")
	go task("manasin wajan")
	go task("nyiapain bumbu")
	go task("masak nasi goreng")
	fmt.Println("nunggu panas wajan dulu")
	time.Sleep(5 * time.Second)

	fmt.Println("main goroutine completed")
}

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v index %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}
