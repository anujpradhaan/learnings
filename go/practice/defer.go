package main

import "log"

func deferTest() {
	log.Println("Testing Defer")
	defer log.Println("I am Defering now")

	log.Println("Inbetween line 1")

	log.Println("Final line") // As there is not return statement here, the defer statement will get executed after the final line in such cases
	//Defer will be executed here
}

func returnAndDefer() int {
	log.Println("Starting in returnAndDefer")
	defer log.Println("I am deferring")

	log.Println("Intermediate line")

	//Defer will get executed here, just before the return statement
	return 1
}

func main() {
	deferTest()
	log.Println(returnAndDefer())
}
