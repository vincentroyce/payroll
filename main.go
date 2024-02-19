package main

import (
	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/payroll/models"
)

func main() {
	uadmin.Register(
		models.Company{},
		models.Employee{},
	)

	uadmin.StartServer()
}
