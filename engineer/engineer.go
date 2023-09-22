package engineer

type Engineer struct {
	position string
	salary   float64
	address  string
}

func NewEngineer(position string, salary float64, address string) *Engineer {
	return &Engineer{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (e *Engineer) GetPosition() string {
	return e.position
}

func (e *Engineer) SetPosition(position string) {
	e.position = position
}

func (e *Engineer) GetSalary() float64 {
	return e.salary
}

func (e *Engineer) SetSalary(salary float64) {
	e.salary = salary
}

func (e *Engineer) GetAddress() string {
	return e.address
}

func (e *Engineer) SetAddress(address string) {
	e.address = address
}
