package main

import (
	"fmt"
	"log"

	service "github.com/caklucky/crud-unit-test/Service"
	"github.com/caklucky/crud-unit-test/config"
	"github.com/caklucky/crud-unit-test/models"
	"github.com/caklucky/crud-unit-test/repository"
	"github.com/google/uuid"
)

// var

func main() {
	fmt.Println("hello world")
	cfg := config.NewConfig()
	db, e := cfg.Connect()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	repo := repository.NewEmployee(db)
	employeeService := service.NewEmployeeService(repo)

	//  insert pegawai
	var pegawai models.Employee
	pegawai.ID = uuid.New().String()
	pegawai.Gender = "M"
	pegawai.Phone = "082234561608"
	pegawai.Name = "Lucky Fernanda R"

	fmt.Println(employeeService.TambahPegawai(pegawai))

	// update pegawai
	pegawai.Name = "Lucky Fernanda RRRR"
	err := employeeService.UpdatePegawai(pegawai)
	if err != nil {
		fmt.Println(err)
	}

	// delete pegawai
	err = employeeService.HapusPegawai(pegawai.ID)
	if err != nil {
		fmt.Println(err)
	}
	// lihat pegawai
	fmt.Println(employeeService.LihatPegawai())
}
