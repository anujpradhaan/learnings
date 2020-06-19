package main

import "log"

type Animal struct {
	name string
}

func callByValue(animal Animal) {
	animal.name = "AnotherName"
}

func callByReference(animal *Animal) {
	animal.name = "anotherName"
}

func main() {
	animal := Animal{
		name: "Name",
	}

	log.Println(animal.name)

	callByValue(animal)
	//Printing After call by value, it should not change the value
	log.Println(animal.name)

	callByReference(&animal)
	//Printing aftter call by reference, it should change the value
	log.Println(animal.name)

}
