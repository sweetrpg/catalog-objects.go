package vo

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core/vo"
)

// Studio value object.
// This value object is a serializable representation of the Studio model.
type StudioVO struct {
	ID         string                 `json:"id" jsonapi:"primary,studio"`
	Name       string                 `json:"name" jsonapi:"attr,name"`
	Website    url.URL                `json:"website" jsonapi:"attr,website"`
	Notes      string                 `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
