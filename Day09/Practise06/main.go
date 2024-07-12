package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// pattern channel
func main() {

	//fanOut()
	//waitForResult()
	//deadLock()
	//readChannel()
	//capacityLengthCh()
}

func deadLock() {
	//construct channel with unbuffering
	ch := make(chan int, 0)

	//goroutine waiting to get value from channel
	a := <-ch
	fmt.Println(a)
}

func closeChannel() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)

	close(ch)

	ch <- 10
	wg.Done()
}

func waitForResult() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()
	d := <-ch
	fmt.Println("parent : received signal :", d)
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// 1. before fanout
func bufferingChannel() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 10
	close(ch)

	fmt.Println("capacity : ", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))

	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
}

func fanOut() {
	children := 2000
	ch := make(chan string, children)
	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(c)
	}
	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : received signal :", children)
	}
	time.Sleep(time.Second)
}

func waitForTask() {
	ch := make(chan string)
	go func() {
		d := <-ch
		fmt.Println("child : received signal :", d)
	}()
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "data"
	fmt.Println("parent : sent signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func pooling() {
	ch := make(chan string)
	g := runtime.GOMAXPROCS(0)
	for c := 0; c < g; c++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d : received signal : %s\n", child, d)
			}
			fmt.Printf("child %d : received shutdown signal\n", child)
		}(c)
	}
	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent : sent signal :", w)
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func drop() {
	const cap = 100
	ch := make(chan string, cap)
	go func() {
		for p := range ch {
			fmt.Println("child : received signal :", p)
		}
	}()
	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", w)
		default:
			fmt.Println("parent : dropped data :", w)
		}
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "data"
	}()
	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func boundedWorkPooling() {
	work := []string{"paper", "rock", "scissor", "donut", 2000: "dunkin"}
	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)
	ch := make(chan string, g)
	for c := 0; c < g; c++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : received signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : received shutdown signal\n", child)
		}(c)
	}
	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

/* 	go func(msg string) {
	fmt.Println(msg)
}("hello")*/
