package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/caklucky/crud-unit-test/models"
)

type EmployeeRepository struct {
	Db *sql.DB
}

const (
	table = "employees"
)

func NewEmployee(db *sql.DB) EmployeeRepository {
	return EmployeeRepository{
		Db: db,
	}
}

func (e *EmployeeRepository) Find() (employees []models.Employee, err error) {

	queryText := fmt.Sprintf("SELECT * FROM %s Order By id DESC", table)

	rowQuery, err := e.Db.Query(queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var employee models.Employee
		if err = rowQuery.Scan(&employee.ID,
			&employee.Name,
			&employee.Gender,
			&employee.Phone,
		); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (e *EmployeeRepository) FindById(id string) (employees models.Employee, err error) {
	queryText := fmt.Sprintf("SELECT * FROM %s where id = '%s'", table, id)
	row := e.Db.QueryRow(queryText)
	err = row.Scan(&employees.ID, &employees.Phone, &employees.Gender,
		&employees.Name)
	return
}

func (e *EmployeeRepository) Insert(employee models.Employee) (err error) {
	queryText := fmt.Sprintf("INSERT INTO %s (id, name, phone, gender) values('%s','%s','%s','%s')", table,
		employee.ID,
		employee.Name,
		employee.Phone,
		employee.Gender,
	)

	_, err = e.Db.Exec(queryText)

	if err != nil {
		return err
	}
	return
}

func (e *EmployeeRepository) Update(employee models.Employee) (err error) {
	fmt.Println(employee)
	queryText := fmt.Sprintf("UPDATE %s set name = %s, phone ='%s', gender = '%s' where id = '%s'",
		table,
		employee.Name,
		employee.Phone,
		employee.Gender,
		employee.ID,
	)

	_, err = e.Db.Exec(queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return
}

func (e *EmployeeRepository) Delete(id string) (err error) {
	queryText := fmt.Sprintf("DELETE FROM %s where id = '%s'", table, id)

	s, err := e.Db.Exec(queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()

	if check == 0 {
		return errors.New("id not found ")
	}

	return
}
