package models

import "github.com/uadmin/uadmin"

type Supervisor struct {
	uadmin.Model
	Employee     Employee
	EmployeeID   uint
	Department   Department
	DepartmentID uint
	Position     Position
	PositionID   uint
	Worksite     Worksite
	WorksiteID   uint
}
