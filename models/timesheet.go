package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Timesheet struct {
	uadmin.Model
	Name       Employee
	NameID     uint
	TimeIn     time.Time
	BreakStart time.Time
	BreakEnd   time.Time
	TimeOut    time.Time
	// Duration (Time-in to Break Start + Break End to Time-out)
}
