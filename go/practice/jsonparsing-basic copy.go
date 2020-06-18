package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//Marshelling bool
	val, _ := json.Marshal(true)
	fmt.Println(val)         // prints just bytes
	fmt.Println(string(val)) //prints the actual string

	//Marshelling int
	val, _ = json.Marshal(1)
	fmt.Println(val)
	fmt.Println(string(val))

	//Marshelling float
	val, _ = json.Marshal(1.1)
	fmt.Println(val)
	fmt.Println(string(val))

	//Marshelling string
	val, _ = json.Marshal("abc")
	fmt.Println(val)
	fmt.Println(string(val))

	//Marshelling string array
	sArray := []string{"a", "1", "3"}
	val, _ = json.Marshal(sArray)
	fmt.Println(val)
	fmt.Println(string(val))

	//Marshelling map
	cmap := map[string]string{
		"name": "anuj", "phone": "9090909090"}
	val, _ = json.Marshal(cmap)
	fmt.Println(val)
	fmt.Println(string(val))
}
