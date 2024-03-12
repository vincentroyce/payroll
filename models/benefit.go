package models

import "github.com/uadmin/uadmin"

type Benefit struct {
	uadmin.Model
	Name       string
	Percentage float64
}
