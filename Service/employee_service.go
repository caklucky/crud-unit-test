package Service

import (
	"errors"

	"github.com/caklucky/crud-unit-test/models"
	"github.com/caklucky/crud-unit-test/repository"
)

type EmployeeService struct {
	Repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
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

func (e *EmployeeService) HapusPegawai(id string) {

}
