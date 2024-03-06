package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) { //needed prefix initialized from main.go
	page := strings.TrimPrefix(r.URL.Path, "/")
	session := uadmin.IsAuthenticated(r)

	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	}

	c := map[string]interface{}{}
	switch page {

	case "dashboard": //Name of HTML
		c = DashboardHandler(w, r)
	case "payroll": //Name of HTML
		c = PayrollHandler(w, r)
	case "employee-list": // Name of HTML
		c = EmployeeListHandler(w, r)
	case "time-sheet": // Name of HTML
		c = TimesheetHandler(w, r)
	case "elements": // Name of HTML
		c = ElementHandler(w, r)
	default:
		page = "dashboard"
	}
	//uadmin.Trail(uadmin.DEBUG, page)
	c["Page"] = page
	uadmin.Trail(uadmin.DEBUG, page)
	// if r.URL.Path == "/logout" {
	// 	/* If the request URL Path is /logout after the /login_system/, it will proceed to this part.
	// 	e.g. localhost:8080/login_system/logout */

	// 	// LogoutHandler handles the logout process for the user.
	// 	LogoutHandler(w, r, session)
	// 	return
	// }
	Rendering(w, r, page, c)

}

func Rendering(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/base.html")

	path := "./templates/" + page + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}

func ElementHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}

	return c
}
