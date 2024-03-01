package models

import "github.com/uadmin/uadmin"

type WorkSites struct {
	uadmin.Model
	Worksite    string
	ManagedBy   Employee
	ManagedByID uint
}
