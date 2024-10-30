package vo

import (
	modelcore "github.com/sweetrpg/model-core/vo"
)

type ContributionVO struct {
	ID     string    `json:"id" jsonapi:"primary,contribution"`
	Person *PersonVO `json:"person,omitempty" jsonapi:"relation,person,omitempty"`
	Volume *VolumeVO `json:"volume,omitempty" jsonapi:"relation,volume,omitempty"`
	Roles  []string  `jsonapi:"attr,roles"`
	Notes  string    `json:"notes" jsonapi:"attr,notes"`
	modelcore.AuditableVO
}
