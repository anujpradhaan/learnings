package main

import (
	"fmt"
	"oops/department"
	"oops/employee"
)

func main() {
	e := employee.Employee{
		FirstName: "Anuj",
		LastName:  "Kumar",
	}
	fmt.Println(e.GetAge())

	// department.GetDepartmentName() -> This won't work as this method belongs to the struct and struct is unexported outside the package itself
	d := department.New("Tech")
	//d := department.Random("Tech") //There is no consturct as New in golang. its just a method which is returning the unexported struct
	fmt.Println(d.GetDepartmentName())
}
