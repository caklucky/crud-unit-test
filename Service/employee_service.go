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
	return employees, err
}

func (e *EmployeeService) TambahPegawai(emp models.Employee) (employee models.Employee, err error) {
	_ = e.Repo.Insert(emp)
	employee, err = e.Repo.FindById(emp.ID)
	return
}

func (e *EmployeeService) UpdatePegawai(emp models.Employee) (employee models.Employee, err error) {
	pegawai, _ := e.Repo.FindById(emp.ID)
	if pegawai == (models.Employee{}) {
		return models.Employee{}, errors.New("id not found ")
	}
	err = e.Repo.Update(emp)
	if err != nil {
		return models.Employee{}, errors.New("update error ")
	}
	employee, err = e.Repo.FindById(emp.ID)
	return
}

func (e *EmployeeService) HapusPegawai(id string) (err error) {
	err = e.Repo.Delete(id)
	return
}
