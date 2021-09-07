package repository

import (
	"github.com/caklucky/crud-unit-test/models"
	"github.com/jinzhu/gorm"
)

type EmployeeRepository struct {
	Db *gorm.DB
}

func NewEmployee() *EmployeeRepository {
	return &EmployeeRepository{}
}

func (e *EmployeeRepository) GetAll() (employees []models.Employee, err error) {
	return
}

func (e *EmployeeRepository) GetById(id string) (employees models.Employee, err error) {
	return
}

func (e *EmployeeRepository) Insert(employee models.Employee) (employes models.Employee, err error) {
	return
}

func (e *EmployeeRepository) Update(employee models.Employee) (employes models.Employee, err error) {
	return
}

func (e *EmployeeRepository) Delete(id string) (employes models.Employee, err error) {
	return
}
