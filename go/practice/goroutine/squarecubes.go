package main

import (
	"fmt"
	"log"
)

func cube(number int, channel chan int) {
	log.Printf("Calculating cube for %d", number)
	channel <- number * number * number
}

func square(number int, channel chan int) {
	log.Printf("Calculating square for %d", number)
	channel <- number * number
}

func output(number int, channel chan int) {
	log.Printf("Calculating output for %d", number)
	channelSquare := make(chan int)
	channelCube := make(chan int)
	go cube(number, channelCube)
	go square(number, channelSquare)
	output := <-channelSquare + <-channelCube
	channel <- output

}

func main() {
	outputChannel := make(chan int)
	go output(2, outputChannel)
	fmt.Println(<-outputChannel)

}
