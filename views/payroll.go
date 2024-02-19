package views

import (
	"net/http"
	"strings"
)

func PayrollHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}

	c["Title"] = "Payroll | Page"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/payroll")
	return c
}
