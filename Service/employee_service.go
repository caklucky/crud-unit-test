package Service

import (
	"errors"

	"github.com/caklucky/crud-unit-test/models"
)

type EmployeeService struct {
	Repo models.EmployeeRepository
}

func NewEmployeeService(repo models.EmployeeRepository) EmployeeService {
	return EmployeeService{
		Repo: repo,
	}
}

func (e *EmployeeService) LihatPegawai() (employees []models.Employee, err error) {
	employees, err = e.Repo.Find()
	return
}

func (e *EmployeeService) TambahPegawai(emp models.Employee) (employee models.Employee, err error) {
	_ = e.Repo.Insert(emp)
	employee, err = e.Repo.FindById(emp.ID)
	return
}

func (e *EmployeeService) UpdatePegawai(employee models.Employee) (err error) {
	pegawai, _ := e.Repo.FindById(employee.ID)
	if pegawai == (models.Employee{}) {
		return errors.New("id not found ")
	}
	err = e.Repo.Update(employee)
	return
}

func (e *EmployeeService) HapusPegawai(id string) (err error) {
	err = e.Repo.Delete(id)
	return
}
