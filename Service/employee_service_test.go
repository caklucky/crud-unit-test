package Service

import (
	"testing"

	"github.com/caklucky/crud-unit-test/models"
	"github.com/caklucky/crud-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var employeeRepository = &repository.EmployeeRepositoryMock{Mock: mock.Mock{}}
var employeeService = NewEmployeeService(employeeRepository)

func TestEmployeeService_LihatPegawaiNotFound(t *testing.T) {

	employeeRepository.Mock.On("Find").Return([]models.Employee{}, nil)

	category, err := employeeService.LihatPegawai()
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
