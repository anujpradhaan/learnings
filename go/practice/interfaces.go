package main

import "fmt"

type Employee interface {
	CalculateSalary() float32
	GetAge() int
}

type Hobby interface {
	GetHobbies() []string
}

type PermanentEmployees struct {
	EmpId    int
	BasicPay float32
	Pf       float32
	Age      int
	Hobbies  []string
}

type ContractEmployees struct {
	EmpId    int
	BasicPay float32
	Age      int
	Hobbies  []string
}

type Freelancer struct {
	EmpId            int
	ratePerHour      float32
	totalHoursWorked int
	Age              int
	Hobbies          []string
}

func (permanent PermanentEmployees) CalculateSalary() float32 {
	return permanent.BasicPay + permanent.Pf
}

func (permanent PermanentEmployees) GetAge() int {
	return permanent.Age
}

func (permanent PermanentEmployees) GetHobbies() []string {
	return permanent.Hobbies
}

func (contract ContractEmployees) CalculateSalary() float32 {
	return contract.BasicPay
}

func (contract ContractEmployees) GetAge() int {
	return contract.Age
}

func (contract ContractEmployees) GetHobbies() []string {
	return contract.Hobbies
}

func (freelancer Freelancer) CalculateSalary() float32 {
	return freelancer.ratePerHour * float32(freelancer.totalHoursWorked)
}

func (freelancer Freelancer) GetAge() int {
	return freelancer.Age
}

func (freelancer Freelancer) GetHobbies() []string {
	return freelancer.Hobbies
}

func calculateTotalSalary(employee []Employee) float32 {
	var expense float32 = 0.0
	for _, value := range employee {
		expense = expense + value.CalculateSalary()
	}
	return expense
}

func calculateAverageAge(employee []Employee) float32 {
	totalAgeSum := 0
	for _, employee := range employee {
		totalAgeSum = employee.GetAge() + totalAgeSum
	}
	return float32(totalAgeSum / len(employee))
}

func getAllHobbies(employeesHobbies []Hobby) []string {
	allHobbies := []string{}
	for _, employee := range employeesHobbies {
		allHobbies = append(allHobbies, getIndividualHobbies(employee)...)
	}
	return allHobbies
}

func getIndividualHobbies(hobby Hobby) []string {
	return hobby.GetHobbies()
}

func main() {
	p1 := PermanentEmployees{
		EmpId:    1,
		BasicPay: 10.5,
		Pf:       1.0,
		Age:      28,
		Hobbies:  []string{"Writing", "Singing"},
	}

	p2 := ContractEmployees{
		EmpId:    2,
		BasicPay: 5.0,
		Age:      30,
		Hobbies:  []string{"Dancing", "Reading"},
	}

	p3 := PermanentEmployees{
		EmpId:    3,
		BasicPay: 12.0,
		Pf:       2.0,
		Age:      56,
		Hobbies:  []string{"Singing", "Music"},
	}

	p4 := Freelancer{
		EmpId:            4,
		totalHoursWorked: 2,
		ratePerHour:      10,
		Age:              21,
		Hobbies:          []string{"coding"},
	}
	employees := []Employee{p1, p2, p3, p4}
	hobbies := []Hobby{p1, p2, p3, p4}
	expense := calculateTotalSalary(employees)
	fmt.Println(expense)
	fmt.Println(calculateAverageAge(employees))
	fmt.Println(getAllHobbies(hobbies))
}
