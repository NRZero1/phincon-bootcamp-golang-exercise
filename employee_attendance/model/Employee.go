package model

type Employee struct {
	Id          int
	Name        string
	Attendances []Attendance
}

type EmployeeBuilder interface {
	SetId(int) EmployeeBuilder
	SetName(string) EmployeeBuilder
	Build() *Employee
}

type employeeBuilder struct {
	employee *Employee
}

func NewEmployeeBuilder() EmployeeBuilder {
	return &employeeBuilder{
		employee: &Employee{},
	}
}

func (eb *employeeBuilder) SetId(id int) EmployeeBuilder {
	eb.employee.Id = id
	return eb
}

func (eb *employeeBuilder) SetName(name string) EmployeeBuilder {
	eb.employee.Name = name
	return eb
}

func (eb *employeeBuilder) Build() *Employee {
	eb.employee.Attendances = []Attendance{}
	return eb.employee
}

func (e *Employee) GetId() int {
	return e.Id
}

func (e *Employee) GetName() string {
	return e.Name
}