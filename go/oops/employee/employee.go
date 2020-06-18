package employee

type Employee struct {
	FirstName string
	LastName  string
	age       int
}

func (e Employee) GetAge() int {
	return e.age
}
