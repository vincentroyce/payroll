package views

import (
	"net/http"
	"strings"
)

// Handler is used to handle/load pages
func TimesheetHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}
	// Initialize Title and mapped it on the html file (you can check it if you want :)
	c["Title"] = "Time Sheet | Page"
	// tells the handler that it will load to this path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/time-sheet")
	return c
}
