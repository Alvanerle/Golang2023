package hr

type HR struct {
	position string
	salary   float64
	address  string
}

func NewHR(position string, salary float64, address string) *HR {
	return &HR{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (h *HR) GetPosition() string {
	return h.position
}

func (h *HR) SetPosition(position string) {
	h.position = position
}

func (h *HR) GetSalary() float64 {
	return h.salary
}

func (h *HR) SetSalary(salary float64) {
	h.salary = salary
}

func (h *HR) GetAddress() string {
	return h.address
}

func (h *HR) SetAddress(address string) {
	h.address = address
}
