package Service

import (
	"testing"

	"github.com/caklucky/crud-unit-test/models"
	"github.com/caklucky/crud-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEmployeeService_LihatPegawaiNotFound(t *testing.T) {
	repo := new(repository.EmployeeRepositoryMock)
	service := NewEmployeeService(repo)
	repo.Mock.On("Find").Return([]models.Employee{}, nil)

	employee, err := service.LihatPegawai()
	assert.Empty(t, employee)
	assert.Nil(t, err)
}

func TestEmployeeService_LihatPegawaiSuccess(t *testing.T) {
	repo := new(repository.EmployeeRepositoryMock)
	service := NewEmployeeService(repo)
	ekspetasi := []models.Employee{
		{
			Name:   "sukijan",
			ID:     "1",
			Gender: "M",
			Phone:  "08112002211",
		},
		{
			Name:   "sukijan 2",
			ID:     "2",
			Gender: "M",
			Phone:  "08112002214",
		},
	}
	repo.Mock.On("Find").Return(ekspetasi, nil)

	result, err := service.LihatPegawai()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Len(t, ekspetasi, len(result))

	assert.Equal(t, ekspetasi[0].ID, result[0].ID)
	assert.Equal(t, ekspetasi[0].Name, result[0].Name)
	repo.AssertCalled(t, "Find")

}

func TestEmployeeService_TambahPegawaiSuccess(t *testing.T) {
	repo := new(repository.EmployeeRepositoryMock)
	service := NewEmployeeService(repo)
	ekspetasi := models.Employee{
		Name:   "sukijan",
		ID:     "1",
		Gender: "M",
		Phone:  "08112002211",
	}
	repo.Mock.On("Insert", mock.AnythingOfType("models.Employee")).Return(nil).Once()
	repo.Mock.On("FindById", mock.AnythingOfType("string")).Return(ekspetasi, nil).Once()

	result, err := service.TambahPegawai(ekspetasi)
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
	assert.Equal(t, ekspetasi.ID, result.ID)
}
