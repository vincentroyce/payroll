package models

import "github.com/uadmin/uadmin"

type Worksite struct {
	uadmin.Model
	Worksite string
}

func (s *Worksite) String() string {
	return s.Worksite
}
