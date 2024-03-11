package api

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/payroll/models"
)

func AddEmployeeAPIHandler(w http.ResponseWriter, r *http.Request) {
	emp := models.Employee{}
	uadminUser := uadmin.User{}

	employeeFName := r.FormValue("fname")
	employeeLName := r.FormValue("lname")
	employeeWorksite := r.FormValue("worksite")
	//employeeDepartment := strconv.Atoi(r.FormValue("department"))
	employeeContact := r.FormValue("contactphone")
	uadmin.Trail(uadmin.DEBUG, employeeFName)

	emp.FirstName = employeeFName
	emp.LastName = employeeLName
	emp.Company.Worksite = employeeWorksite
	//emp.Department = employeeDepartment
	emp.ContactNumber = employeeContact

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
