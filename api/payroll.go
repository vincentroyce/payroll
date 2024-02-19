package api

import (
	"net/http"
	"strings"
)

function PayrollAPIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path := strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
}