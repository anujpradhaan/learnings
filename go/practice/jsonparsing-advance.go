package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name   string
	Phone  []string
	Gender string
	Age    int
}

type person1 struct {
	Name   string   `json:"name"`
	Phones []string `json:"phones"`
	Gender string   `json:"gender"`
	Age    int      `json:"age"`
}

type person2 struct {
	Name   string   `json:"name"`
	Phones []string `json:"phones"`
	Gender string   `json:"gender"`
}

func main() {
	//Marshelling a type struct
	p := &person{Name: "Anuj", Phone: []string{"90909090909", "545454545"}, Gender: "Male", Age: 31}
	v, _ := json.Marshal(p)
	fmt.Println(string(v))

	ap := &person1{Name: "Ashvik", Phones: []string{"dsdsfdsf", "dsfdfsdff"}, Gender: "M", Age: 1}
	v, _ = json.Marshal(ap)
	fmt.Println(string(v))

	//Unmarshelling the same values
	p1 := person1{}
	json.Unmarshal(v, &p1) // you can still pass non pointer value, but it will throw an error at runtime  -> json: Unmarshal(non-pointer main.person1)
	fmt.Println(p1)

	//What if we convert the same value to s differnt data structure with less number of fields
	//In that case, it will skip the values missing in the destination data stucture.
	p2 := person2{}
	json.Unmarshal(v, &p2) //
	fmt.Println(p2)

	/*
		The above Unmarshelling will work when we have a defined data structure to which we want to store data.
		What if we don't always know the data coming in the incoming json.
	*/
	var data map[string]interface{}
	json.Unmarshal(v, &data) // ti will result in map where values are coming in sorted order of key names.
	//Also, By default the numerical values come as float64
	fmt.Println(data["age"].(float64) == 1)

	//If we are actually looking for type safety then defining a data structure with apropriate datatypes will give you less
	//error prone results

}
