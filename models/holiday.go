package models

import "github.com/uadmin/uadmin"

type Holiday struct {
	uadmin.Model
	Name  string
	Hours string
	Rate  int //
}
