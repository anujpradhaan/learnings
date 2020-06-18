package classes

type Employee struct {
	FirstName string
	LastName  string
	age       int
}

func (e Employee) getAge() int {
	return e.age
}
