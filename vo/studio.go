package vo

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// Studio value object.
// This value object is a serializable representation of the Studio model.
//
// Website is a plain string, not url.URL - see PublisherVO's comment.
type StudioVO struct {
	ID         string                 `json:"id" jsonapi:"primary,studio"`
	Name       string                 `json:"name" jsonapi:"attr,name"`
	Website    string                 `json:"website" jsonapi:"attr,website"`
	Notes      string                 `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	// Audit fields are direct fields, not an embedded modelcore.AuditableVO: google/jsonapi
	// v1.0.0 skips untagged struct fields, so an embedded struct never reaches the wire.
	// Times are RFC 3339 strings, which is what catalog-web's decoder expects.
	CreatedAt time.Time  `json:"created_at" jsonapi:"attr,created_at,rfc3339"`
	CreatedBy string     `json:"created_by" jsonapi:"attr,created_by"`
	UpdatedAt time.Time  `json:"updated_at" jsonapi:"attr,updated_at,rfc3339"`
	UpdatedBy string     `json:"updated_by" jsonapi:"attr,updated_by"`
	DeletedAt *time.Time `json:"deleted_at" jsonapi:"attr,deleted_at,omitempty,rfc3339"`
	DeletedBy *string    `json:"deleted_by" jsonapi:"attr,deleted_by,omitempty"`
}
