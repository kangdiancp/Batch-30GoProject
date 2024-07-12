package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chanelSendObject()
	//fanOut()
	//waitForResult()
	//deadLock()
	//readChannel()
	//capacityLengthCh()
}

func fanOut() {
	children := 2000
	ch := make(chan string, children)
	for c := 0; c < children; c++ {
		go func(child int) { //define parameter
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(c) //passing value via argument
	}
	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : received signal :", children)
	}
	time.Sleep(time.Second)
}

type Account struct {
	saldo float64
}

func chanelSendObject() {
	ch := make(chan Account)
	fmt.Println(ch)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- Account{saldo: 1_000}
		fmt.Println("child : sent signal :")
	}()
	d := <-ch
	fmt.Println("parent(main) goroutine : received signal :", d)
	time.Sleep(time.Second)

}

func waitForResult() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()
	d := <-ch
	fmt.Println("parent(main) goroutine : received signal :", d)
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func deadLock() {
	ch := make(chan int, 3)

	//a coba baca ch
	a := <-ch
	fmt.Println("a :", a)
}

func readChannel() {
	//1.construct channel
	ch := make(chan int, 3)

	//2. send data to chanel
	ch <- 5
	ch <- 10
	ch <- 15

	//closing channel
	close(ch)

	n, open := <-ch
	fmt.Printf("Received: %d, open : %t\n", n, open)

	n, open = <-ch
	fmt.Printf("Received: %d, open : %t\n", n, open)

	n, open = <-ch
	fmt.Printf("Received: %d, open : %t\n", n, open)

	n, open = <-ch
	fmt.Printf("Received: %d, open : %t\n", n, open) //get error

}

func capacityLengthCh() {
	//1.construct channel
	ch := make(chan int, 3)

	//2. send data to chanel
	ch <- 5
	ch <- 10
	ch <- 15

	close(ch)

	fmt.Println("capacity:", cap(ch))
	fmt.Println("length : ", len(ch))

	fmt.Println("read value : ", <-ch) //goroutine read ch#1
	fmt.Println("new length is :", len(ch))

	fmt.Println("read value : ", <-ch) //goroutine read ch#2
	fmt.Println("new length is :", len(ch))
}
