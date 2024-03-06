package api

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/payroll/models"
)

func LoginAPIHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	uadmin.Trail(uadmin.DEBUG, "username api: ", username)
	uadmin.Trail(uadmin.DEBUG, "password api: ", password)
	login := []models.Employee{}
	uadmin.Filter(&login, "username = ? AND password = ?", username, password)
	for i := range login {
		uadmin.Preload(&login[i])
	}
	uadmin.ReturnJSON(w, r, login)
}
