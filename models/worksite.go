package models

import "github.com/uadmin/uadmin"

type WorkSite struct {
	uadmin.Model
	Worksite string
}

func (s *WorkSite) String() string {
	return s.Worksite
}
