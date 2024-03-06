package models

import "github.com/uadmin/uadmin"

type WorkSite struct {
	uadmin.Model
	Worksite    string
	ManagedBy   Employee
	ManagedByID uint
}
