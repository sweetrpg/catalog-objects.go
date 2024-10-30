package vo

import (
	modelcore "github.com/sweetrpg/model-core/vo"
)

// Person value object.
// This value object is a serializable representation of the Person model.
type PersonVO struct {
	ID         string                 `json:"id" jsonapi:"primary,person"`
	Name       string                 `json:"name" jsonapi:"attr,name"`
	Notes      string                 `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
