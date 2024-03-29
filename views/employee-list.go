package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/payroll/models"
)

// Handler is used to handle/load pages
func EmployeeListHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}
	employee := []models.Employee{}
	department := []models.Department{}
	position := []models.Position{}
	worksite := []models.Worksite{}
	// Initialize Title and mapped it on the html file (you can check it if you want :)

	uadmin.All(&employee)
	for x := range employee {
		uadmin.Preload(&employee[x])
	}
	c["Title"] = "Employee List | Page"
	c["Employee"] = employee

	uadmin.All(&department)
	for x := range department {
		uadmin.Preload(&employee[x])
	}
	c["Department"] = department

	uadmin.All(&position)
	for x := range position {
		uadmin.Preload(&position[x])
	}
	c["Position"] = position

	uadmin.All(&worksite)
	for x := range worksite {
		uadmin.Preload(&worksite[x])
	}
	c["Worksite"] = worksite
	// tells the handler that it will load to this path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/employee-list")
	return c
}
