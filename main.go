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
	uadmin.RootURL = "/admin/"
	http.HandleFunc("/", uadmin.Handler(views.MainHandler)) 
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))

	uadmin.StartServer()
}
