package views

import (
	"net/http"
	"strings"
)

function PayrollViewsHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path := strings.TrimPrefix(r.URL.Path, "/payroll")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
}