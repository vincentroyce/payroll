package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Timesheet struct {
	uadmin.Model
	Username   string
	Name       Employee
	NameID     uint
	TimeIn     time.Time
	BreakStart time.Time `uadmin:"list_exclude"`
	BreakEnd   time.Time `uadmin:"list_exclude"`
	TimeOut    time.Time
	// Duration (Time-in to Break Start + Break End to Time-out)
}
