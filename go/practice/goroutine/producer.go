package main

import "log"

func producer(channel chan int) {
	log.Println("Producing values from 0-9 and pushing them to the input channel")
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func main() {
	channel := make(chan int)
	go producer(channel)

	// infinite loop to test closing of channel
	for {
		value, ok := <-channel
		if ok == false {
			log.Println("Channel closed")
			break
		}
		log.Println(value)
	}

	log.Println("Doing again")

	//Channel once closed can't even be used and there won't be any method calls even if we pass that channel as an argument to any goroutine
	channel = make(chan int)
	go producer(channel)

	for v := range channel {
		log.Println(v)
	}
}
