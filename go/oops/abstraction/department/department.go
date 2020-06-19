package department

type department struct {
	name string
}

func (d department) GetDepartmentName() string {
	return d.name
}

func New(n string) department {
	return department{
		name: n,
	}
}
