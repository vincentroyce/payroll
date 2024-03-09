package models

import "github.com/uadmin/uadmin"

type Department struct {
	uadmin.Model
	Name       string
	WorkSite   WorkSite
	WorkSiteID uint
	// IsCostCenter string
	// IsRevenueCenter string
	// IsDirectCost string
}
