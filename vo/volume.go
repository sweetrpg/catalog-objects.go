package vo

import (
	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// Volume value object.
// This value object is a serializable representation of the Volume model.
type VolumeVO struct {
	ID          string                 `json:"id" jsonapi:"primary,volume"`
	Title       string                 `json:"title" jsonapi:"attr,title"`
	Description string                 `json:"description" jsonapi:"attr,description"`
	Notes       string                 `json:"notes" jsonapi:"attr,notes"`
	Systems     []SystemVO             `json:"systems" jsonapi:"relation,system"`
	Publishers  []PublisherVO          `json:"publishers" jsonapi:"relation,publisher"`
	Studios     []StudioVO             `json:"studios" jsonapi:"relation,studio"`
	Licenses    []LicenseVO            `json:"licenses" jsonapi:"relation,license"`
	Properties  []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags        []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
