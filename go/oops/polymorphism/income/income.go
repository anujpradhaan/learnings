package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

func (f FixedBilling) calculate() int {
	return f.biddedAmount
}

func (f FixedBilling) source() string {
	return f.projectName
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func getCostToTheCompany(incomes []Income) int {
	cost := 0
	for _, income := range incomes {
		cost += income.calculate()
	}
	return cost
}

func main() {
	f1 := FixedBilling{
		projectName:  "Project1",
		biddedAmount: 20,
	}

	f2 := FixedBilling{
		projectName:  "Project2",
		biddedAmount: 30,
	}

	tm1 := TimeAndMaterial{
		projectName: "Project3",
		noOfHours:   10,
		hourlyRate:  2,
	}

	tm2 := TimeAndMaterial{
		projectName: "Project3",
		noOfHours:   30,
		hourlyRate:  3,
	}

	income := []Income{f1, f2, tm1, tm2}
	fmt.Println(getCostToTheCompany(income))
}
