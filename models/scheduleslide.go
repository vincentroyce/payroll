package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type ScheduleSlide struct {
	uadmin.Model
	Name                   string
	FromDate               time.Time
	ToDate                 time.Time
	NewClockIn             time.Time
	NewClockout            time.Time
	ApprovedBySupervisor   Employee
	ApprovedBySupervisorID uint
	ApprovedByDepartment   Department
	ApprovedByDepartmentID uint
}
