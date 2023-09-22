package manager

type Manager struct {
	position string
	salary   float64
	address  string
}

func NewManager(position string, salary float64, address string) *Manager {
	return &Manager{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) GetSalary() float64 {
	return m.salary
}

func (m *Manager) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}
