package Service

import (
	"fmt"
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

	employee, err := employeeService.LihatPegawai()
	assert.Empty(t, employee)
	assert.Nil(t, err)
}

func TestEmployeeService_LihatPegawaiSuccess(t *testing.T) {
	ekspetasi := []models.Employee{{
		Name:   "sukijan",
		ID:     "1",
		Gender: "M",
		Phone:  "08112002211",
	},
	}
	employeeRepository.Mock.On("Find").Return(ekspetasi, nil)

	result, err := employeeService.LihatPegawai()
	fmt.Println("test -->", result)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, ekspetasi[0].ID, result[0].ID)
	assert.Equal(t, ekspetasi[0].Name, result[0].Name)
}

func TestEmployeeService_TambahPegawaiSuccess(t *testing.T) {
	ekspetasi := []models.Employee{{
		Name:   "sukijan",
		ID:     "1",
		Gender: "M",
		Phone:  "08112002211",
	},
	}
	employeeRepository.Mock.On("FindById", "1").Return(ekspetasi, nil)

	result, err := employeeService.LihatPegawai()
	fmt.Println("test -->", result)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, ekspetasi[0].ID, result[0].ID)
	assert.Equal(t, ekspetasi[0].Name, result[0].Name)
}
