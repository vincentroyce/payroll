package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/payroll/api"
	"github.com/vincentroyce/payroll/models"
	"github.com/vincentroyce/payroll/views"
)

func main() {
	uadmin.Register(
		//models.Company{},
		models.Employee{},
		models.Allowances{},
		models.Benefit{},
		models.Deduction{},
		models.Department{},
		models.Holiday{},
		models.Leaves{},
		models.Offset{},
		models.Overtime{},
		models.ScheduleSlide{},
		models.Timesheet{},
		models.WorkSite{},
		models.Position{},
	)
	uadmin.RootURL = "/admin/"
	http.HandleFunc("/api/delete-employee/", uadmin.Handler(api.DeleteEmployeeAPIHandler))
	http.HandleFunc("/api/dashboard/", uadmin.Handler(api.LoginAPIHandler))
	http.HandleFunc("/api/add-employee/", uadmin.Handler(api.AddEmployeeAPIHandler))
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))

	uadmin.StartServer()
}
