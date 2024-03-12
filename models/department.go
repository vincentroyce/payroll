package models

import "github.com/uadmin/uadmin"

type Department struct {
	uadmin.Model
	Name       string
	Worksite   Worksite
	WorksiteID uint

	// IsCostCenter string
	// IsRevenueCenter string
	// IsDirectCost string
}
