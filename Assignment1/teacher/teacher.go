package teacher

type Teacher struct {
	position string
	salary   float64
	address  string
}

func NewTeacher(position string, salary float64, address string) *Teacher {
	return &Teacher{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (t *Teacher) GetPosition() string {
	return t.position
}

func (t *Teacher) SetPosition(position string) {
	t.position = position
}

func (t *Teacher) GetSalary() float64 {
	return t.salary
}

func (t *Teacher) SetSalary(salary float64) {
	t.salary = salary
}

func (t *Teacher) GetAddress() string {
	return t.address
}

func (t *Teacher) SetAddress(address string) {
	t.address = address
}
