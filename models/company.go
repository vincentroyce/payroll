package models

import "github.com/uadmin/uadmin"

type Company struct {
	uadmin.Model
	Logo        string `uadmin:"image"`
	Name        string
	Description string `uadmin:"html"`
	Address     string
}
