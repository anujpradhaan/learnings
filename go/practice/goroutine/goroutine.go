package main

import (
	"fmt"
	"time"
)

var a chan int
var b chan int

func HelloA() {
	fmt.Println("Inside Hello GoRoutine", <-a)
}

func HelloB() {
	fmt.Println("Inside Hello GoRoutine", <-b)
}

func main() {
	a = make(chan int)
	b = make(chan int)
	go HelloA()
	a <- 1
	time.Sleep(1 * time.Second)
	go HelloB()
	b <- 2
	time.Sleep(1 * time.Second)

}
