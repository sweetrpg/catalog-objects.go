package vo

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core/vo"
)

// Publisher value object.
// This value object is a serializable representation of the Publisher model.
type PublisherVO struct {
	ID         string                 `json:"id" jsonapi:"primary,publisher"`
	Name       string                 `json:"name" jsonapi:"attr,name"`
	Address    string                 `json:"address" jsonapi:"attr,address"`
	Website    url.URL                `json:"website" jsonapi:"attr,website"`
	Notes      string                 `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
