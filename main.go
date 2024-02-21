package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/payroll/models"
	"github.com/vincentroyce/payroll/views"
)

func main() {
	uadmin.Register(
		models.Company{},
		models.Employee{},
	)

	http.HandleFunc("/system/", uadmin.Handler(views.MainHandler)) // Set the page you want to show first

	uadmin.StartServer()
}
