package models

type Employee struct {
	ID     string
	Name   string
	Gender string
	Phone  string
}

type EmployeeRepository interface {
	Find() (employees []Employee, err error)
	FindById(id string) (employees Employee, err error)
	Insert(employee Employee) (err error)
	Update(employee Employee) (err error)
	Delete(id string) (err error)
}
