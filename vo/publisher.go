package vo

import (
	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// Publisher value object.
// This value object is a serializable representation of the Publisher model.
//
// Website is a plain string, not url.URL - url.URL has no MarshalJSON, so jsonapi previously
// serialized it as its internal struct fields (Scheme/Host/Path/...) instead of a URL string,
// breaking every non-Go client.
type PublisherVO struct {
	ID         string                 `json:"id" jsonapi:"primary,publisher"`
	Name       string                 `json:"name" jsonapi:"attr,name"`
	Address    string                 `json:"address" jsonapi:"attr,address"`
	Website    string                 `json:"website" jsonapi:"attr,website"`
	Notes      string                 `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
