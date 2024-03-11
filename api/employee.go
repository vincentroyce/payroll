package api

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/payroll/models"
)

func DeleteEmployeeAPIHandler(w http.ResponseWriter, r *http.Request) {
	employee := []models.Employee{}
	uadminUsers := []uadmin.User{}

	ids := strings.Split(r.FormValue("ID"), ",")

	uadmin.Trail(uadmin.DEBUG, "IDs: %v", ids)

	uadmin.Filter(&uadminUsers, "username IN ?", ids)
	for x := range uadminUsers {
		uadminUsers[x].Active = false
		uadminUsers[x].Save()
	}

	uadmin.DeleteList(&employee, "id_number IN ?", ids)
	for i := range employee {
		employee[i].Save()
	}
}
