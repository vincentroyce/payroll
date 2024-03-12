package api

import (
	"net/http"
	"strconv"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/payroll/models"
)

func AddEmployeeAPIHandler(w http.ResponseWriter, r *http.Request) {
	emp := models.Employee{}
	uadminUser := uadmin.User{}

	employeeFName := r.FormValue("fname")
	employeeLName := r.FormValue("lname")
	employeeWorksite, _ := strconv.ParseUint(r.FormValue("worksite"), 10, 64)
	employeeDepartment, _ := strconv.ParseUint(r.FormValue("department"), 10, 64)
	employeeContact := r.FormValue("contactphone")
	uadmin.Trail(uadmin.DEBUG, employeeFName)

	emp.FirstName = employeeFName
	emp.LastName = employeeLName
	emp.WorksiteID = uint(employeeWorksite)
	emp.DepartmentID = uint(employeeDepartment)
	emp.ContactNumber = employeeContact
	uadmin.Trail(uadmin.DEBUG, "worksite :", employeeDepartment)
	uadmin.Trail(uadmin.DEBUG, "worksite :", emp.Department)
	uadminUser.FirstName = employeeFName
	uadminUser.LastName = employeeLName
	uadminUser.Username = employeeFName
	uadminUser.Password = "Password11"
	uadminUser.Active = true
	emp.IdNumber = employeeFName

	//uadmin.Trail(uadmin.DEBUG, employeeFName)
	uadminUser.Save()
	emp.Save()

	// if err != nil {
	// 	// Return an error message if the database did not save properly.
	// 	uadmin.ReturnJSON(w, r, map[string]interface{}{
	// 		"status":  "error",
	// 		"err_msg": "Error saving the database : " + err.Error(),
	// 	})
	// 	return
	// }

	// // Return JSON to the client.
	// uadmin.ReturnJSON(w, r, map[string]interface{}{
	// 	"status": "ok",
	// })
}
