package models

import "github.com/uadmin/uadmin"

type Department struct {
	uadmin.Model
	WorkSites   WorkSites
	WorkSitesID uint
	Managedby   WorkSites
	ManagedbyID uint
	// IsCostCenter string
	// IsRevenueCenter string
	// IsDirectCost string
}
