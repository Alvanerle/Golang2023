package sales

type Sales struct {
	position string
	salary   float64
	address  string
}

func NewSales(position string, salary float64, address string) *Sales {
	return &Sales{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (s *Sales) GetPosition() string {
	return s.position
}

func (s *Sales) SetPosition(position string) {
	s.position = position
}

func (s *Sales) GetSalary() float64 {
	return s.salary
}

func (s *Sales) SetSalary(salary float64) {
	s.salary = salary
}

func (s *Sales) GetAddress() string {
	return s.address
}

func (s *Sales) SetAddress(address string) {
	s.address = address
}
