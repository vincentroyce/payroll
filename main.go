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
		models.Worksite{},
		models.Position{},
		models.Supervisor{},
	)
	// uadmin.RegisterInlines(
	// 	models.Employee{},
	// 	map[string]string{
	// 		"Department": "EmployeeID",
	// 	},
	// )
	uadmin.RegisterInlines(
		models.Employee{},
		map[string]string{
			"Supervisor": "EmployeeID",
		},
	)
	uadmin.RegisterInlines(
		models.Department{},
		map[string]string{
			"Employee": "DepartmentID",
		},
	)
	uadmin.RegisterInlines(
		models.Position{},
		map[string]string{
			"Employee":   "PositionID",
			"Supervisor": "PositionID",
		},
	)
	uadmin.RegisterInlines(
		models.Worksite{},
		map[string]string{
			"Department": "WorksiteID",
			"Supervisor": "WorksiteID",
		},
	)

	// uadmin.RegisterInlines(
	// 	models.Employee{},
	// 	map[string]string{
	// 		"Department": "EmployeeID",
	// 	})
	uadmin.RootURL = "/admin/"
	http.HandleFunc("/api/delete-employee/", uadmin.Handler(api.DeleteEmployeeAPIHandler))
	http.HandleFunc("/api/dashboard/", uadmin.Handler(api.LoginAPIHandler))
	http.HandleFunc("/api/add-employee/", uadmin.Handler(api.AddEmployeeAPIHandler))
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))

	uadmin.StartServer()
}
